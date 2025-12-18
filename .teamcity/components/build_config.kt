/*
 * Copyright IBM Corp. 2021, 2025
 * SPDX-License-Identifier: MPL-2.0
 */

import jetbrains.buildServer.configs.kotlin.BuildSteps
import jetbrains.buildServer.configs.kotlin.buildSteps.ScriptBuildStep
import jetbrains.buildServer.configs.kotlin.buildSteps.script
import java.io.File

fun BuildSteps.ConfigureGoEnv() {
    step(ScriptBuildStep {
        name = "Configure GOENV"
        scriptContent = File("./scripts/configure_goenv.sh").readText()
    })
}