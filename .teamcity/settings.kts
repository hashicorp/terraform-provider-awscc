import jetbrains.buildServer.configs.kotlin.* // ktlint-disable no-wildcard-imports
import jetbrains.buildServer.configs.kotlin.buildFeatures.freeDiskSpace
import jetbrains.buildServer.configs.kotlin.buildFeatures.golang
import jetbrains.buildServer.configs.kotlin.buildSteps.script
import jetbrains.buildServer.configs.kotlin.failureConditions.failOnText
import jetbrains.buildServer.configs.kotlin.failureConditions.BuildFailureOnText
import jetbrains.buildServer.configs.kotlin.triggers.schedule
import java.io.File
import java.time.Duration
import java.time.LocalTime
import java.time.ZoneId
import java.time.format.DateTimeFormatter

version = "2023.05"

val defaultRegion = DslContext.getParameter("default_region")
val awsAccountID = DslContext.getParameter("aws_account.account_id")
val acctestParallelism = DslContext.getParameter("acctest_parallelism", "")
val tfAccAssumeRoleArn = DslContext.getParameter("tf_acc_assume_role_arn", "")
val tfLog = DslContext.getParameter("tf_log", "")

// Assume Role credentials
val accTestRoleARN = DslContext.getParameter("aws_account.role_arn", "")
val awsAccessKeyID = if (accTestRoleARN != "") { DslContext.getParameter("aws_account.access_key_id") } else { "" }
val awsSecretAccessKey = if (accTestRoleARN != "") { DslContext.getParameter("aws_account.secret_access_key") } else { "" }

project {
    if (DslContext.getParameter("build_full", "true").toBoolean()) {
        buildType(FullBuild)
    }

    params {
        if (acctestParallelism != "") {
            text("ACCTEST_PARALLELISM", acctestParallelism, allowEmpty = false)
        }
        text("TEST_PATTERN", "TestAcc", display = ParameterDisplay.HIDDEN)
        text("env.AWS_ACCOUNT_ID", awsAccountID, display = ParameterDisplay.HIDDEN, allowEmpty = false)
        text("env.AWS_DEFAULT_REGION", defaultRegion, allowEmpty = false)
        text("env.TF_LOG", tfLog)

        val brancRef = DslContext.getParameter("branch_name", "")
        if (brancRef != "") {
            text("BRANCH_NAME", brancRef, display = ParameterDisplay.HIDDEN)
        }

        if (tfAccAssumeRoleArn != "") {
            text("env.TF_ACC_ASSUME_ROLE_ARN", tfAccAssumeRoleArn)
        }

        // Assume Role credentials
        password("AWS_ACCESS_KEY_ID", awsAccessKeyID, display = ParameterDisplay.HIDDEN)
        password("AWS_SECRET_ACCESS_KEY", awsSecretAccessKey, display = ParameterDisplay.HIDDEN)
        text("ACCTEST_ROLE_ARN", accTestRoleARN, display = ParameterDisplay.HIDDEN)

        // Define this parameter even when not set to allow individual builds to set the value
        text("env.TF_ACC_TERRAFORM_VERSION", DslContext.getParameter("terraform_version", ""))

        // These overrides exist because of the inherited dependency in the existing project structure and can
        // be removed when this is moved outside of it
        val isOnPrem = DslContext.getParameter("is_on_prem", "true").equals("true", ignoreCase = true)
        if (isOnPrem) {
            // These should be overridden in the base AWS project
            param("env.GOPATH", "")
            param("env.GO111MODULE", "") // No longer needed as of Go 1.16
            param("env.GO_VERSION", "") // We're using `goenv` and `.go-version`
        }
    }
}

object FullBuild : BuildType({
    name = "Acceptance Tests - Main"

    vcs {
        root(AbsoluteId(DslContext.getParameter("vcs_root_id")))

        cleanCheckout = true
    }

    steps {
        ConfigureGoEnv()
        script {
            name = "Run Acceptance Tests"
            scriptContent = File("./scripts/acceptance_tests.sh").readText()
        }
    }

    val runNightly = DslContext.getParameter("run_nightly_build", "")
    if (runNightly.toBoolean()) {
        val triggerTimeRaw = DslContext.getParameter("trigger_time")
        val formatter = DateTimeFormatter.ofPattern("HH':'mm' 'VV")
        val triggerTime = formatter.parse(triggerTimeRaw)
        val triggerDay = if (DslContext.getParameter("trigger_day", "") != "") {
            DslContext.getParameter("trigger_day", "")
        } else {
            "Sun-Thu"
        }

        val enableTestTriggersGlobally = DslContext.getParameter("enable_test_triggers_globally", "true").equals("true", ignoreCase = true)
        if (enableTestTriggersGlobally) {
            triggers {
                schedule {
                    schedulingPolicy = cron {
                        dayOfWeek = triggerDay
                        val triggerHM = LocalTime.from(triggerTime)
                        hours = triggerHM.getHour().toString()
                        minutes = triggerHM.getMinute().toString()
                        timezone = ZoneId.from(triggerTime).toString()
                    }
                    branchFilter = "+:refs/heads/main"
                    triggerBuild = always()
                    withPendingChangesOnly = false
                    enableQueueOptimization = false
                    enforceCleanCheckoutForDependencies = true
                }
            }
        }
    }

    features {
        feature {
            type = "JetBrains.SharedResources"
            param("locks-param", "${DslContext.getParameter("aws_account.lock_id")} writeLock")
        }

        val freeDiskSpace = DslContext.getParameter("free_disk_space", "")
        if (freeDiskSpace.toBoolean()) {
            freeDiskSpace {
                failBuild = true
            }
        }
    }
})