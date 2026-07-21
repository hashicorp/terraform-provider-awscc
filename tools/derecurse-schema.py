# Copyright IBM Corp. 2021, 2026
# SPDX-License-Identifier: MPL-2.0

"""Break recursive $ref cycles in a CloudFormation resource type schema.

Terraform's type system cannot model recursive attribute definitions
(https://github.com/hashicorp/terraform-provider-awscc/issues/95), so the
code generator never terminates on schemas whose definitions reference
themselves (e.g. AWS::WAFv2::WebACL's Statement -> AndStatement -> Statement).

This script rewrites such a schema in place, unrolling every recursion cycle
to a fixed depth (default 3) - the same trade-off terraform-provider-aws makes
in its hand-written WAFv2 schemas:

1. Recursion cycles among "definitions" are detected generically, as strongly
   connected components of the $ref graph (no type names are hardcoded).
2. Each cycle has "counting" members that every cycle passes through (for
   WAFv2 this is "Statement"); each $ref into a counting member increments
   the depth level. Cycle members are cloned once per level: level 1 keeps
   the original name, deeper clones get a suffix (Statement, StatementLevel2,
   StatementLevel3).
3. A clone whose *required* properties would exceed the maximum depth is
   unconstructible (e.g. a level-3 AndStatement, since its required
   Statements items would be level-4); *optional* properties that reference
   an out-of-depth or unconstructible definition are dropped. So a level-3
   Statement keeps RateBasedStatement/ManagedRuleGroupStatement (minus their
   optional ScopeDownStatement) but drops AndStatement/OrStatement/
   NotStatement - exactly how terraform-provider-aws models it.

As a preliminary, semantically neutral normalization, $refs to pure alias
definitions (definitions whose body is exactly {"$ref": ...}, e.g. WebACL's
AddressField -> FieldIdentifier) are rewritten to point at the alias target;
the code generator cannot type array items that resolve through such aliases
(https://github.com/hashicorp/terraform-provider-awscc/issues/1515).

The script is deterministic (sorted iteration) and idempotent: a schema with
no aliases and no remaining cycles is left unchanged.

Usage:
    python3 tools/derecurse-schema.py [--depth N] <schema.json> [<schema.json> ...]
"""

import argparse
import json
import re
import sys
from copy import deepcopy

REF = re.compile(r"^#/definitions/([^/]+)$")


def collect_refs(node, acc):
    """Append all local definition names referenced from node to acc."""
    if isinstance(node, dict):
        for key in sorted(node):
            value = node[key]
            if key == "$ref" and isinstance(value, str):
                match = REF.match(value)
                if match:
                    acc.append(match.group(1))
            else:
                collect_refs(value, acc)
    elif isinstance(node, list):
        for value in node:
            collect_refs(value, acc)
    return acc


def strongly_connected_components(graph):
    """Iterative Tarjan SCC over graph (dict: node -> sorted neighbors).

    Returns cyclic components only (size > 1, or a self-referencing node),
    sorted for determinism.
    """
    index, low = {}, {}
    on_stack, stack = set(), []
    counter = [0]
    components = []

    for root in sorted(graph):
        if root in index:
            continue
        index[root] = low[root] = counter[0]
        counter[0] += 1
        stack.append(root)
        on_stack.add(root)
        work = [(root, iter(graph[root]))]
        while work:
            node, neighbors = work[-1]
            advanced = False
            for neighbor in neighbors:
                if neighbor not in graph:
                    continue
                if neighbor not in index:
                    index[neighbor] = low[neighbor] = counter[0]
                    counter[0] += 1
                    stack.append(neighbor)
                    on_stack.add(neighbor)
                    work.append((neighbor, iter(graph[neighbor])))
                    advanced = True
                    break
                if neighbor in on_stack:
                    low[node] = min(low[node], index[neighbor])
            if advanced:
                continue
            work.pop()
            if work:
                parent = work[-1][0]
                low[parent] = min(low[parent], low[node])
            if low[node] == index[node]:
                component = []
                while True:
                    member = stack.pop()
                    on_stack.discard(member)
                    component.append(member)
                    if member == node:
                        break
                if len(component) > 1 or node in graph[node]:
                    components.append(sorted(component))
    return sorted(components)


def is_cyclic(nodes, graph):
    """True if the subgraph induced by nodes contains a cycle."""
    nodes = set(nodes)
    color = {}  # 1 = visiting, 2 = done

    for root in sorted(nodes):
        if color.get(root) == 2:
            continue
        color[root] = 1
        work = [(root, iter(graph[root]))]
        while work:
            node, neighbors = work[-1]
            advanced = False
            for neighbor in neighbors:
                if neighbor not in nodes:
                    continue
                state = color.get(neighbor)
                if state == 1:
                    return True
                if state is None:
                    color[neighbor] = 1
                    work.append((neighbor, iter(graph[neighbor])))
                    advanced = True
                    break
            if not advanced:
                color[node] = 2
                work.pop()
    return False


def counting_members(component, graph):
    """Members that lie on every cycle of the component.

    Removing such a member leaves the component acyclic, so every pass around
    any cycle goes through it; counting depth on refs into these members makes
    the unroll depth measure nesting levels of the recursive type itself
    (levels of Statement for WAFv2). If no single member covers all cycles,
    every member counts (plain unrolling).
    """
    members = set(component)
    counting = {m for m in component if not is_cyclic(members - {m}, graph)}
    return counting or members


def clone_name(name, level):
    return name if level == 1 else f"{name}Level{level}"


class Unroller:
    """Unrolls one recursion cycle to a fixed depth."""

    def __init__(self, definitions, component, graph, depth):
        self.definitions = definitions
        self.members = set(component)
        self.counting = counting_members(component, graph)
        self.depth = depth
        self.pruned_properties = set()

    def target_level(self, target, level):
        """Level of the clone that a ref to `target` from level `level` uses."""
        return level + 1 if target in self.counting else level

    def refs_ok(self, node, level, constructible):
        """True if all in-cycle refs under node stay in depth and constructible."""
        for target in collect_refs(node, []):
            if target not in self.members:
                continue
            target_level = self.target_level(target, level)
            if target_level > self.depth:
                return False
            if not constructible.get((target, target_level), False):
                return False
        return True

    def compute_constructible(self):
        """Fixed point: which (member, level) clones can exist at all.

        A clone is unconstructible if a *required* property, or any part of
        the definition outside "properties", needs an out-of-depth or
        unconstructible clone.
        """
        constructible = {
            (member, level): True
            for member in self.members
            for level in range(1, self.depth + 1)
        }
        changed = True
        while changed:
            changed = False
            for (member, level), ok in sorted(constructible.items()):
                if not ok:
                    continue
                body = self.definitions[member]
                properties = body.get("properties", {})
                required = set(body.get("required", []))
                for key in sorted(body):
                    node = properties if key == "properties" else body[key]
                    if key == "properties":
                        bad = any(
                            not self.refs_ok(properties[p], level, constructible)
                            for p in required
                            if p in properties
                        )
                    else:
                        bad = not self.refs_ok(node, level, constructible)
                    if bad:
                        constructible[(member, level)] = False
                        changed = True
                        break
        return constructible

    def build_clone(self, member, level, constructible):
        """Clone `member` at `level`: drop out-of-depth optional properties,
        redirect surviving in-cycle refs to the right clone names."""
        body = deepcopy(self.definitions[member])
        properties = body.get("properties")
        if isinstance(properties, dict):
            for prop in sorted(properties):
                if not self.refs_ok(properties[prop], level, constructible):
                    del properties[prop]
                    self.pruned_properties.add(prop)
        self.redirect(body, level)
        return body

    def redirect(self, node, level):
        """Point in-cycle $refs under node at the clone for their level."""
        if isinstance(node, dict):
            for key in sorted(node):
                value = node[key]
                if key == "$ref" and isinstance(value, str):
                    match = REF.match(value)
                    if match and match.group(1) in self.members:
                        target = match.group(1)
                        target_level = self.target_level(target, level)
                        node[key] = f"#/definitions/{clone_name(target, target_level)}"
                else:
                    self.redirect(value, level)
        elif isinstance(node, list):
            for value in node:
                self.redirect(value, level)

    def unroll(self):
        constructible = self.compute_constructible()

        for member in sorted(self.members):
            if not constructible[(member, 1)]:
                raise SystemExit(
                    f"error: {member} is unconstructible even at level 1; "
                    f"increase --depth"
                )
            for level in range(2, self.depth + 1):
                name = clone_name(member, level)
                if name in self.definitions:
                    raise SystemExit(f"error: definition {name} already exists")

        # Build clones reachable from the level-1 (original) definitions.
        clones = {}
        pending = {(member, 1) for member in self.members}
        while pending:
            member, level = min(pending)
            pending.discard((member, level))
            if (member, level) in clones:
                continue
            clone = self.build_clone(member, level, constructible)
            clones[(member, level)] = clone
            for target in collect_refs(clone, []):
                base = re.sub(r"Level(\d+)$", "", target)
                match = re.search(r"Level(\d+)$", target)
                if base in self.members:
                    target_level = int(match.group(1)) if match else 1
                    if (base, target_level) not in clones:
                        pending.add((base, target_level))

        for member in sorted(self.members):
            self.definitions[member] = clones[(member, 1)]
        for member, level in sorted(clones):
            if level > 1:
                self.definitions[clone_name(member, level)] = clones[(member, level)]

        return (
            f"cycle {{{', '.join(sorted(self.members))}}} unrolled to depth "
            f"{self.depth} (counting refs into {', '.join(sorted(self.counting))}; "
            f"{len(clones) - len(self.members)} cloned definitions added)"
        )


def inline_aliases(schema):
    """Rewrite $refs to pure alias definitions to their targets.

    An alias definition's body is exactly {"$ref": "#/definitions/X"}. The
    code generator cannot type array items that resolve through an alias, so
    point every ref at the alias's (transitively resolved) target instead.
    The now-unreferenced alias definitions are kept; they are harmless.
    Returns a list of change notes.
    """
    definitions = schema.get("definitions")
    if not isinstance(definitions, dict):
        return []

    aliases = {}
    for name in sorted(definitions):
        body = definitions[name]
        if isinstance(body, dict) and set(body) == {"$ref"}:
            match = REF.match(body["$ref"])
            if match:
                aliases[name] = match.group(1)

    def resolve(name, seen=()):
        if name in seen:
            raise SystemExit(f"error: alias cycle at {name}")
        return resolve(aliases[name], seen + (name,)) if name in aliases else name

    targets = {name: resolve(name) for name in aliases}

    inlined = set()

    def rewrite(node, inside_alias):
        if isinstance(node, dict):
            for key in sorted(node):
                value = node[key]
                if (
                    key == "$ref"
                    and isinstance(value, str)
                    and not inside_alias
                ):
                    match = REF.match(value)
                    if match and match.group(1) in targets:
                        node[key] = f"#/definitions/{targets[match.group(1)]}"
                        inlined.add(match.group(1))
                else:
                    rewrite(value, inside_alias)
        elif isinstance(node, list):
            for value in node:
                rewrite(value, inside_alias)

    for key in sorted(schema):
        if key == "definitions":
            for name in sorted(definitions):
                rewrite(definitions[name], name in aliases)
        else:
            rewrite(schema[key], False)

    return [
        f"alias {name} inlined to {targets[name]}" for name in sorted(inlined)
    ]


def derecurse(schema, depth):
    """Unroll all recursion cycles in schema; returns a list of change notes.

    Also records the set of property names pruned at the depth boundary under
    the "x-derecursed" key so the code generator can surface truncation at
    runtime (a warning when an imported/read resource carries them beyond the
    modeled depth).
    """
    definitions = schema.get("definitions")
    if not isinstance(definitions, dict):
        return []

    graph = {
        name: sorted(set(collect_refs(definitions[name], [])))
        for name in definitions
    }
    notes = []
    pruned = set()
    for component in strongly_connected_components(graph):
        unroller = Unroller(definitions, component, graph, depth)
        notes.append(unroller.unroll())
        pruned |= unroller.pruned_properties
    if pruned:
        schema["x-derecursed"] = {
            "depth": depth,
            "prunedProperties": sorted(pruned),
        }
    return notes


def main():
    parser = argparse.ArgumentParser(
        description="Unroll recursive $ref cycles in CloudFormation schemas."
    )
    parser.add_argument(
        "--depth",
        type=int,
        default=3,
        help="maximum nesting depth for recursive types (default: 3)",
    )
    parser.add_argument("schemas", nargs="+", metavar="schema.json")
    args = parser.parse_args()

    if args.depth < 1:
        parser.error("--depth must be >= 1")

    for path in args.schemas:
        with open(path, encoding="utf-8") as fh:
            schema = json.load(fh)
        notes = inline_aliases(schema)
        notes += derecurse(schema, args.depth)
        if not notes:
            print(f"{path}: no aliases or recursion cycles found")
            continue
        with open(path, "w", encoding="utf-8") as fh:
            json.dump(schema, fh, indent=2)
            fh.write("\n")
        for note in notes:
            print(f"{path}: {note}")


if __name__ == "__main__":
    sys.exit(main())
