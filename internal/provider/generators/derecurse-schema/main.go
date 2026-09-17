// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

// derecurse-schema breaks recursive $ref cycles in a CloudFormation resource type schema.
//
// Terraform's type system cannot model recursive attribute definitions
// (https://github.com/hashicorp/terraform-provider-awscc/issues/95), so the
// code generator never terminates on schemas whose definitions reference
// themselves (e.g. AWS::WAFv2::WebACL's Statement -> AndStatement -> Statement).
//
// This tool rewrites such a schema in place, unrolling every recursion cycle
// to a fixed depth (default 3) - the same trade-off terraform-provider-aws makes
// in its hand-written WAFv2 schemas:
//
//  1. Recursion cycles among "definitions" are detected generically, as strongly
//     connected components of the $ref graph (no type names are hardcoded).
//  2. Each cycle has "counting" members that every cycle passes through (for
//     WAFv2 this is "Statement"); each $ref into a counting member increments
//     the depth level. Cycle members are cloned once per level: level 1 keeps
//     the original name, deeper clones get a suffix (Statement, StatementLevel2,
//     StatementLevel3).
//  3. A clone whose *required* properties would exceed the maximum depth is
//     unconstructible (e.g. a level-3 AndStatement, since its required
//     Statements items would be level-4); *optional* properties that reference
//     an out-of-depth or unconstructible definition are dropped. So a level-3
//     Statement keeps RateBasedStatement/ManagedRuleGroupStatement (minus their
//     optional ScopeDownStatement) but drops AndStatement/OrStatement/
//     NotStatement - exactly how terraform-provider-aws models it.
//
// As a preliminary, semantically neutral normalization, $refs to pure alias
// definitions (definitions whose body is exactly {"$ref": ...}, e.g. WebACL's
// AddressField -> FieldIdentifier) are rewritten to point at the alias target;
// the code generator cannot type array items that resolve through such aliases
// (https://github.com/hashicorp/terraform-provider-awscc/issues/1515).
//
// The tool is deterministic (sorted iteration) and idempotent: a schema with
// no aliases and no remaining cycles is left unchanged. The set of property
// names pruned at the depth boundary is recorded under the schema's
// "x-derecursed" key so the code generator can surface truncation at runtime.
//
// Usage:
//
//	go run internal/provider/generators/derecurse-schema/main.go [--depth N] <schema.json> [<schema.json> ...]
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/common"
)

const (
	defaultDepth = 3 // terraform-provider-aws's hand-written WAFv2 trade-off
	indentStep   = 2 // matches Python's json.dump(..., indent=2)
)

var (
	refPattern   = regexp.MustCompile(`^#/definitions/([^/]+)$`)
	levelPattern = regexp.MustCompile(`Level(\d+)$`)
)

// object is a JSON object that preserves key order, so that a rewritten
// schema serializes with its original key order (as Python's json module
// does) and untouched parts of the file don't churn.
type object struct {
	keys   []string
	values map[string]any
}

func newObject() *object {
	return &object{values: make(map[string]any)}
}

func (o *object) get(key string) (any, bool) {
	v, ok := o.values[key]
	return v, ok
}

func (o *object) has(key string) bool {
	_, ok := o.values[key]
	return ok
}

// set replaces the value in place if the key exists, else appends it.
func (o *object) set(key string, value any) {
	if !o.has(key) {
		o.keys = append(o.keys, key)
	}
	o.values[key] = value
}

func (o *object) delete(key string) {
	if !o.has(key) {
		return
	}
	delete(o.values, key)
	o.keys = slices.DeleteFunc(o.keys, func(k string) bool { return k == key })
}

func (o *object) sortedKeys() []string {
	keys := slices.Clone(o.keys)
	sort.Strings(keys)
	return keys
}

// parseValue decodes the next JSON value from dec into *object / []any /
// string / json.Number / bool / nil, preserving object key order.
func parseValue(dec *json.Decoder) (any, error) {
	token, err := dec.Token()
	if err != nil {
		return nil, err
	}

	switch token := token.(type) {
	case json.Delim:
		switch token {
		case '{':
			obj := newObject()
			for dec.More() {
				keyToken, err := dec.Token()
				if err != nil {
					return nil, err
				}
				key, ok := keyToken.(string)
				if !ok {
					return nil, fmt.Errorf("expected object key, got %v", keyToken)
				}
				value, err := parseValue(dec)
				if err != nil {
					return nil, err
				}
				obj.set(key, value)
			}
			if _, err := dec.Token(); err != nil { // consume '}'
				return nil, err
			}
			return obj, nil
		case '[':
			arr := []any{}
			for dec.More() {
				value, err := parseValue(dec)
				if err != nil {
					return nil, err
				}
				arr = append(arr, value)
			}
			if _, err := dec.Token(); err != nil { // consume ']'
				return nil, err
			}
			return arr, nil
		default:
			return nil, fmt.Errorf("unexpected delimiter %v", token)
		}
	default:
		return token, nil
	}
}

func parseJSON(data []byte) (any, error) {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	value, err := parseValue(dec)
	if err != nil {
		return nil, err
	}
	if _, err := dec.Token(); err != io.EOF {
		return nil, fmt.Errorf("trailing data after JSON value")
	}
	return value, nil
}

// writeJSONString escapes s the way Python's json module does with its
// default ensure_ascii=True: printable ASCII stays literal, everything
// else becomes an escape sequence.
func writeJSONString(buf *bytes.Buffer, s string) {
	buf.WriteByte('"')
	for _, r := range s {
		switch r {
		case '"':
			buf.WriteString(`\"`)
		case '\\':
			buf.WriteString(`\\`)
		case '\n':
			buf.WriteString(`\n`)
		case '\r':
			buf.WriteString(`\r`)
		case '\t':
			buf.WriteString(`\t`)
		case '\b':
			buf.WriteString(`\b`)
		case '\f':
			buf.WriteString(`\f`)
		default:
			switch {
			case r >= 0x20 && r <= 0x7e:
				buf.WriteRune(r)
			case r > 0xffff: //nolint:mnd // BMP boundary: encode as a UTF-16 surrogate pair
				r -= 0x10000
				fmt.Fprintf(buf, `\u%04x\u%04x`, 0xd800+(r>>10), 0xdc00+(r&0x3ff)) //nolint:mnd
			default:
				fmt.Fprintf(buf, `\u%04x`, r)
			}
		}
	}
	buf.WriteByte('"')
}

// writeJSON serializes value with 2-space indentation, matching Python's
// json.dump(value, fh, indent=2) byte for byte.
func writeJSON(buf *bytes.Buffer, value any, indent int) {
	switch value := value.(type) {
	case *object:
		if len(value.keys) == 0 {
			buf.WriteString("{}")
			return
		}
		buf.WriteString("{\n")
		for i, key := range value.keys {
			buf.WriteString(strings.Repeat(" ", indent+indentStep))
			writeJSONString(buf, key)
			buf.WriteString(": ")
			writeJSON(buf, value.values[key], indent+indentStep)
			if i < len(value.keys)-1 {
				buf.WriteByte(',')
			}
			buf.WriteByte('\n')
		}
		buf.WriteString(strings.Repeat(" ", indent))
		buf.WriteByte('}')
	case []any:
		if len(value) == 0 {
			buf.WriteString("[]")
			return
		}
		buf.WriteString("[\n")
		for i, item := range value {
			buf.WriteString(strings.Repeat(" ", indent+indentStep))
			writeJSON(buf, item, indent+indentStep)
			if i < len(value)-1 {
				buf.WriteByte(',')
			}
			buf.WriteByte('\n')
		}
		buf.WriteString(strings.Repeat(" ", indent))
		buf.WriteByte(']')
	case string:
		writeJSONString(buf, value)
	case json.Number:
		buf.WriteString(value.String())
	case int:
		buf.WriteString(strconv.Itoa(value))
	case bool:
		buf.WriteString(strconv.FormatBool(value))
	case nil:
		buf.WriteString("null")
	default:
		panic(fmt.Sprintf("unexpected JSON value type %T", value))
	}
}

func deepCopy(value any) any {
	switch value := value.(type) {
	case *object:
		clone := newObject()
		for _, key := range value.keys {
			clone.set(key, deepCopy(value.values[key]))
		}
		return clone
	case []any:
		clone := make([]any, len(value))
		for i, item := range value {
			clone[i] = deepCopy(item)
		}
		return clone
	default:
		return value
	}
}

// collectRefs appends all local definition names referenced from node to acc.
func collectRefs(node any, acc []string) []string {
	switch node := node.(type) {
	case *object:
		for _, key := range node.sortedKeys() {
			value := node.values[key]
			if key == "$ref" {
				if ref, ok := value.(string); ok {
					if m := refPattern.FindStringSubmatch(ref); m != nil {
						acc = append(acc, m[1])
						continue
					}
				}
			}
			acc = collectRefs(value, acc)
		}
	case []any:
		for _, value := range node {
			acc = collectRefs(value, acc)
		}
	}
	return acc
}

// stronglyConnectedComponents runs Tarjan's SCC algorithm over graph
// (node -> sorted neighbors). It returns cyclic components only (size > 1,
// or a self-referencing node), each sorted, in sorted order for determinism.
func stronglyConnectedComponents(graph map[string][]string) [][]string {
	index := make(map[string]int)
	low := make(map[string]int)
	onStack := make(map[string]bool)
	var stack []string
	counter := 0
	var components [][]string

	type frame struct {
		node string
		next int
	}

	roots := make([]string, 0, len(graph))
	for node := range graph {
		roots = append(roots, node)
	}
	sort.Strings(roots)

	for _, root := range roots {
		if _, seen := index[root]; seen {
			continue
		}
		index[root], low[root] = counter, counter
		counter++
		stack = append(stack, root)
		onStack[root] = true
		work := []frame{{node: root}}
		for len(work) > 0 {
			top := &work[len(work)-1]
			node := top.node
			advanced := false
			for top.next < len(graph[node]) {
				neighbor := graph[node][top.next]
				top.next++
				if _, inGraph := graph[neighbor]; !inGraph {
					continue
				}
				if _, seen := index[neighbor]; !seen {
					index[neighbor], low[neighbor] = counter, counter
					counter++
					stack = append(stack, neighbor)
					onStack[neighbor] = true
					work = append(work, frame{node: neighbor})
					advanced = true
					break
				}
				if onStack[neighbor] {
					low[node] = min(low[node], index[neighbor])
				}
			}
			if advanced {
				continue
			}
			work = work[:len(work)-1]
			if len(work) > 0 {
				parent := work[len(work)-1].node
				low[parent] = min(low[parent], low[node])
			}
			if low[node] == index[node] {
				var component []string
				for {
					member := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					delete(onStack, member)
					component = append(component, member)
					if member == node {
						break
					}
				}
				if len(component) > 1 || slices.Contains(graph[node], node) {
					sort.Strings(component)
					components = append(components, component)
				}
			}
		}
	}

	slices.SortFunc(components, slices.Compare)
	return components
}

// isCyclic reports whether the subgraph induced by nodes contains a cycle.
func isCyclic(nodes map[string]bool, graph map[string][]string) bool {
	const (
		visiting = 1
		done     = 2
	)
	color := make(map[string]int)

	var visit func(node string) bool
	visit = func(node string) bool {
		color[node] = visiting
		for _, neighbor := range graph[node] {
			if !nodes[neighbor] {
				continue
			}
			switch color[neighbor] {
			case visiting:
				return true
			case 0:
				if visit(neighbor) {
					return true
				}
			}
		}
		color[node] = done
		return false
	}

	roots := make([]string, 0, len(nodes))
	for node := range nodes {
		roots = append(roots, node)
	}
	sort.Strings(roots)
	for _, root := range roots {
		if color[root] == 0 && visit(root) {
			return true
		}
	}
	return false
}

// countingMembers returns the members that lie on every cycle of the component.
//
// Removing such a member leaves the component acyclic, so every pass around
// any cycle goes through it; counting depth on refs into these members makes
// the unroll depth measure nesting levels of the recursive type itself
// (levels of Statement for WAFv2). If no single member covers all cycles,
// every member counts (plain unrolling).
func countingMembers(component []string, graph map[string][]string) map[string]bool {
	members := make(map[string]bool, len(component))
	for _, member := range component {
		members[member] = true
	}
	counting := make(map[string]bool)
	for _, member := range component {
		without := make(map[string]bool, len(members))
		for m := range members {
			without[m] = m != member
		}
		delete(without, member)
		if !isCyclic(without, graph) {
			counting[member] = true
		}
	}
	if len(counting) == 0 {
		return members
	}
	return counting
}

func cloneName(name string, level int) string {
	if level == 1 {
		return name
	}
	return fmt.Sprintf("%sLevel%d", name, level)
}

type memberLevel struct {
	member string
	level  int
}

func (ml memberLevel) compare(other memberLevel) int {
	if c := strings.Compare(ml.member, other.member); c != 0 {
		return c
	}
	return ml.level - other.level
}

// unroller unrolls one recursion cycle to a fixed depth.
type unroller struct {
	definitions *object
	members     map[string]bool
	counting    map[string]bool
	depth       int
	pruned      map[string]bool
}

func newUnroller(definitions *object, component []string, graph map[string][]string, depth int) *unroller {
	members := make(map[string]bool, len(component))
	for _, member := range component {
		members[member] = true
	}
	return &unroller{
		definitions: definitions,
		members:     members,
		counting:    countingMembers(component, graph),
		depth:       depth,
		pruned:      make(map[string]bool),
	}
}

// targetLevel is the level of the clone that a ref to target from level uses.
func (u *unroller) targetLevel(target string, level int) int {
	if u.counting[target] {
		return level + 1
	}
	return level
}

// refsOK reports whether all in-cycle refs under node stay in depth and constructible.
func (u *unroller) refsOK(node any, level int, constructible map[memberLevel]bool) bool {
	for _, target := range collectRefs(node, nil) {
		if !u.members[target] {
			continue
		}
		targetLevel := u.targetLevel(target, level)
		if targetLevel > u.depth {
			return false
		}
		if !constructible[memberLevel{target, targetLevel}] {
			return false
		}
	}
	return true
}

// computeConstructible finds the fixed point of which (member, level) clones
// can exist at all.
//
// A clone is unconstructible if a *required* property, or any part of the
// definition outside "properties", needs an out-of-depth or unconstructible
// clone.
func (u *unroller) computeConstructible() map[memberLevel]bool {
	constructible := make(map[memberLevel]bool)
	for member := range u.members {
		for level := 1; level <= u.depth; level++ {
			constructible[memberLevel{member, level}] = true
		}
	}

	keys := make([]memberLevel, 0, len(constructible))
	for key := range constructible {
		keys = append(keys, key)
	}
	slices.SortFunc(keys, memberLevel.compare)

	for changed := true; changed; {
		changed = false
		for _, key := range keys {
			if !constructible[key] {
				continue
			}
			bodyValue, _ := u.definitions.get(key.member)
			body, ok := bodyValue.(*object)
			if !ok {
				continue
			}
			properties, _ := body.get("properties")
			required := make(map[string]bool)
			if requiredValue, ok := body.get("required"); ok {
				if items, ok := requiredValue.([]any); ok {
					for _, item := range items {
						if name, ok := item.(string); ok {
							required[name] = true
						}
					}
				}
			}
			for _, bodyKey := range body.sortedKeys() {
				var bad bool
				if bodyKey == "properties" {
					if propertiesObject, ok := properties.(*object); ok {
						for name := range required {
							if value, ok := propertiesObject.get(name); ok && !u.refsOK(value, key.level, constructible) {
								bad = true
								break
							}
						}
					}
				} else {
					value, _ := body.get(bodyKey)
					bad = !u.refsOK(value, key.level, constructible)
				}
				if bad {
					constructible[key] = false
					changed = true
					break
				}
			}
		}
	}
	return constructible
}

// buildClone clones member at level: drops out-of-depth optional properties
// and redirects surviving in-cycle refs to the right clone names.
func (u *unroller) buildClone(member string, level int, constructible map[memberLevel]bool) any {
	original, _ := u.definitions.get(member)
	body := deepCopy(original)
	if bodyObject, ok := body.(*object); ok {
		if propertiesValue, ok := bodyObject.get("properties"); ok {
			if properties, ok := propertiesValue.(*object); ok {
				for _, name := range properties.sortedKeys() {
					value, _ := properties.get(name)
					if !u.refsOK(value, level, constructible) {
						properties.delete(name)
						u.pruned[name] = true
					}
				}
			}
		}
	}
	u.redirect(body, level)
	return body
}

// redirect points in-cycle $refs under node at the clone for their level.
func (u *unroller) redirect(node any, level int) {
	switch node := node.(type) {
	case *object:
		for _, key := range node.sortedKeys() {
			value := node.values[key]
			if key == "$ref" {
				if ref, ok := value.(string); ok {
					if m := refPattern.FindStringSubmatch(ref); m != nil && u.members[m[1]] {
						target := m[1]
						node.values[key] = "#/definitions/" + cloneName(target, u.targetLevel(target, level))
						continue
					}
				}
			}
			u.redirect(value, level)
		}
	case []any:
		for _, value := range node {
			u.redirect(value, level)
		}
	}
}

func (u *unroller) sortedMembers() []string {
	members := make([]string, 0, len(u.members))
	for member := range u.members {
		members = append(members, member)
	}
	sort.Strings(members)
	return members
}

func (u *unroller) unroll() (string, error) {
	constructible := u.computeConstructible()

	for _, member := range u.sortedMembers() {
		if !constructible[memberLevel{member, 1}] {
			return "", fmt.Errorf("error: %s is unconstructible even at level 1; increase --depth", member)
		}
		for level := 2; level <= u.depth; level++ {
			name := cloneName(member, level)
			if u.definitions.has(name) {
				return "", fmt.Errorf("error: definition %s already exists", name)
			}
		}
	}

	// Build clones reachable from the level-1 (original) definitions.
	clones := make(map[memberLevel]any)
	pending := make(map[memberLevel]bool)
	for member := range u.members {
		pending[memberLevel{member, 1}] = true
	}
	for len(pending) > 0 {
		var next memberLevel
		first := true
		for key := range pending {
			if first || key.compare(next) < 0 {
				next = key
				first = false
			}
		}
		delete(pending, next)
		if _, done := clones[next]; done {
			continue
		}
		clone := u.buildClone(next.member, next.level, constructible)
		clones[next] = clone
		for _, target := range collectRefs(clone, nil) {
			base := levelPattern.ReplaceAllString(target, "")
			if !u.members[base] {
				continue
			}
			targetLevel := 1
			if m := levelPattern.FindStringSubmatch(target); m != nil {
				targetLevel, _ = strconv.Atoi(m[1])
			}
			if _, done := clones[memberLevel{base, targetLevel}]; !done {
				pending[memberLevel{base, targetLevel}] = true
			}
		}
	}

	for _, member := range u.sortedMembers() {
		u.definitions.set(member, clones[memberLevel{member, 1}])
	}
	cloneKeys := make([]memberLevel, 0, len(clones))
	for key := range clones {
		cloneKeys = append(cloneKeys, key)
	}
	slices.SortFunc(cloneKeys, memberLevel.compare)
	for _, key := range cloneKeys {
		if key.level > 1 {
			u.definitions.set(cloneName(key.member, key.level), clones[key])
		}
	}

	counting := make([]string, 0, len(u.counting))
	for member := range u.counting {
		counting = append(counting, member)
	}
	sort.Strings(counting)

	return fmt.Sprintf("cycle {%s} unrolled to depth %d (counting refs into %s; %d cloned definitions added)",
		strings.Join(u.sortedMembers(), ", "), u.depth, strings.Join(counting, ", "), len(clones)-len(u.members)), nil
}

// inlineAliases rewrites $refs to pure alias definitions to their targets.
//
// An alias definition's body is exactly {"$ref": "#/definitions/X"}. The
// code generator cannot type array items that resolve through an alias, so
// point every ref at the alias's (transitively resolved) target instead.
// The now-unreferenced alias definitions are kept; they are harmless.
// Returns a list of change notes.
func inlineAliases(schema *object) ([]string, error) {
	definitionsValue, _ := schema.get("definitions")
	definitions, ok := definitionsValue.(*object)
	if !ok {
		return nil, nil
	}

	aliases := make(map[string]string)
	for _, name := range definitions.sortedKeys() {
		body, ok := definitions.values[name].(*object)
		if !ok || len(body.keys) != 1 || !body.has("$ref") {
			continue
		}
		if ref, ok := body.values["$ref"].(string); ok {
			if m := refPattern.FindStringSubmatch(ref); m != nil {
				aliases[name] = m[1]
			}
		}
	}

	var resolve func(name string, seen []string) (string, error)
	resolve = func(name string, seen []string) (string, error) {
		if slices.Contains(seen, name) {
			return "", fmt.Errorf("error: alias cycle at %s", name)
		}
		target, isAlias := aliases[name]
		if !isAlias {
			return name, nil
		}
		return resolve(target, append(seen, name))
	}

	// Resolve in sorted order so that a malformed schema with an alias cycle
	// always reports the same member.
	aliasNames := make([]string, 0, len(aliases))
	for name := range aliases {
		aliasNames = append(aliasNames, name)
	}
	sort.Strings(aliasNames)

	targets := make(map[string]string, len(aliases))
	for _, name := range aliasNames {
		target, err := resolve(name, nil)
		if err != nil {
			return nil, err
		}
		targets[name] = target
	}

	inlined := make(map[string]bool)

	var rewrite func(node any, insideAlias bool)
	rewrite = func(node any, insideAlias bool) {
		switch node := node.(type) {
		case *object:
			for _, key := range node.sortedKeys() {
				value := node.values[key]
				if key == "$ref" && !insideAlias {
					if ref, ok := value.(string); ok {
						if m := refPattern.FindStringSubmatch(ref); m != nil {
							if target, isAlias := targets[m[1]]; isAlias {
								node.values[key] = "#/definitions/" + target
								inlined[m[1]] = true
								continue
							}
						}
					}
				}
				rewrite(value, insideAlias)
			}
		case []any:
			for _, value := range node {
				rewrite(value, insideAlias)
			}
		}
	}

	for _, key := range schema.sortedKeys() {
		if key == "definitions" {
			for _, name := range definitions.sortedKeys() {
				_, isAlias := aliases[name]
				rewrite(definitions.values[name], isAlias)
			}
		} else {
			rewrite(schema.values[key], false)
		}
	}

	inlinedNames := make([]string, 0, len(inlined))
	for name := range inlined {
		inlinedNames = append(inlinedNames, name)
	}
	sort.Strings(inlinedNames)

	notes := make([]string, 0, len(inlinedNames))
	for _, name := range inlinedNames {
		notes = append(notes, fmt.Sprintf("alias %s inlined to %s", name, targets[name]))
	}
	return notes, nil
}

// derecurse unrolls all recursion cycles in schema; it returns a list of
// change notes.
//
// It also records the set of property names pruned at the depth boundary
// under the "x-derecursed" key so the code generator can surface truncation
// at runtime (a warning when an imported/read resource carries them beyond
// the modeled depth).
func derecurse(schema *object, depth int) ([]string, error) {
	definitionsValue, _ := schema.get("definitions")
	definitions, ok := definitionsValue.(*object)
	if !ok {
		return nil, nil
	}

	graph := make(map[string][]string, len(definitions.keys))
	for _, name := range definitions.keys {
		refs := collectRefs(definitions.values[name], nil)
		sort.Strings(refs)
		graph[name] = slices.Compact(refs)
	}

	var notes []string
	pruned := make(map[string]bool)
	for _, component := range stronglyConnectedComponents(graph) {
		u := newUnroller(definitions, component, graph, depth)
		note, err := u.unroll()
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
		for name := range u.pruned {
			pruned[name] = true
		}
	}

	if len(pruned) > 0 {
		prunedNames := make([]string, 0, len(pruned))
		for name := range pruned {
			prunedNames = append(prunedNames, name)
		}
		sort.Strings(prunedNames)
		prunedProperties := make([]any, 0, len(prunedNames))
		for _, name := range prunedNames {
			prunedProperties = append(prunedProperties, name)
		}
		marker := newObject()
		marker.set("depth", depth)
		marker.set("prunedProperties", prunedProperties)
		schema.set("x-derecursed", marker)
	}
	return notes, nil
}

// processSchema rewrites the schema in data (if it has aliases or recursion
// cycles) and returns the rewritten bytes (nil if unchanged) and change notes.
func processSchema(data []byte, depth int) ([]byte, []string, error) {
	value, err := parseJSON(data)
	if err != nil {
		return nil, nil, err
	}
	schema, ok := value.(*object)
	if !ok {
		return nil, nil, fmt.Errorf("schema is not a JSON object")
	}

	notes, err := inlineAliases(schema)
	if err != nil {
		return nil, nil, err
	}
	moreNotes, err := derecurse(schema, depth)
	if err != nil {
		return nil, nil, err
	}
	notes = append(notes, moreNotes...)

	if len(notes) == 0 {
		return nil, nil, nil
	}

	var buf bytes.Buffer
	writeJSON(&buf, schema, 0)
	buf.WriteByte('\n')
	return buf.Bytes(), notes, nil
}

func usage() {
	fmt.Fprintf(os.Stderr, "Unroll recursive $ref cycles in CloudFormation schemas.\n\n")
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "\tmain.go [flags] <schema.json> [<schema.json> ...]\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	depth := flag.Int("depth", defaultDepth, "maximum nesting depth for recursive types")
	flag.Usage = usage
	flag.Parse()

	if *depth < 1 {
		fmt.Fprintln(os.Stderr, "error: --depth must be >= 1")
		os.Exit(2)
	}
	paths := flag.Args()
	if len(paths) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	g := common.NewGenerator()

	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			g.Fatalf("error: %s", err)
		}
		rewritten, notes, err := processSchema(data, *depth)
		if err != nil {
			g.Fatalf("%s", err)
		}
		if rewritten == nil {
			g.Infof("%s: no aliases or recursion cycles found", path)
			continue
		}
		if err := os.WriteFile(path, rewritten, 0644); err != nil { //nolint:mnd
			g.Fatalf("error: %s", err)
		}
		for _, note := range notes {
			g.Infof("%s: %s", path, note)
		}
	}
}
