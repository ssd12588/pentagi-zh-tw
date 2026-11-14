package processor

// Docker operations messages
const (
	MsgPullingImage                = "正在拉取映像：%s"
	MsgImagePullCompleted          = "已完成拉取 %s"
	MsgImagePullFailed             = "拉取映像 %s 失敗：%v"
	MsgRemovingWorkerContainers    = "正在移除 worker 容器"
	MsgStoppingContainer           = "正在停止容器 %s"
	MsgRemovingContainer           = "正在移除容器 %s"
	MsgContainerRemoved            = "已移除容器 %s"
	MsgNoWorkerContainersFound     = "找不到 worker 容器"
	MsgWorkerContainersRemoved     = "已移除 %d 個 worker 容器"
	MsgRemovingImage               = "正在移除映像：%s"
	MsgImageRemoved                = "已成功移除映像 %s"
	MsgImageNotFound               = "找不到映像 %s（可能已移除）"
	MsgWorkerImagesRemoveCompleted = "worker 映像移除完成"
	MsgEnsuringDockerNetworks      = "正在確保 Docker 網路存在"
	MsgDockerNetworkExists         = "Docker 網路已存在：%s"
	MsgCreatingDockerNetwork       = "正在建立 Docker 網路：%s"
	MsgDockerNetworkCreated        = "已建立 Docker 網路：%s"
	MsgDockerNetworkCreateFailed   = "建立 Docker 網路 %s 失敗：%v"
	MsgRecreatingDockerNetwork     = "正在以 compose 標籤重建 Docker 網路：%s"
	MsgDockerNetworkRemoved        = "已移除 Docker 網路：%s"
	MsgDockerNetworkRemoveFailed   = "移除 Docker 網路 %s 失敗：%v"
	MsgDockerNetworkInUse          = "Docker 網路 %s 正被容器使用，無法重新建立"
)

// File system operations messages
const (
	MsgExtractingDockerCompose          = "正在解出 docker-compose.yml"
	MsgExtractingLangfuseCompose        = "正在解出 docker-compose-langfuse.yml"
	MsgExtractingObservabilityCompose   = "正在解出 docker-compose-observability.yml"
	MsgExtractingObservabilityDirectory = "正在解出可觀測性目錄"
	MsgSkippingExternalLangfuse         = "略過外部 Langfuse 部署"
	MsgSkippingExternalObservability    = "略過外部可觀測性部署"
	MsgPatchingComposeFile              = "正在修補 docker-compose 檔案：%s"
	MsgComposePatchCompleted            = "docker-compose 檔案修補完成"
	MsgCleaningUpStackFiles             = "正在清理 %s 的堆疊檔案"
	MsgStackFilesCleanupCompleted       = "堆疊檔案清理完成"
	MsgEnsurngStackIntegrity            = "正在確保 %s 堆疊完整性"
	MsgVerifyingStackIntegrity          = "正在驗證 %s 堆疊完整性"
	MsgStackIntegrityVerified           = "%s 堆疊完整性驗證通過"
	MsgUpdatingExistingFile             = "正在更新既有檔案：%s"
	MsgCreatingMissingFile              = "正在建立缺少的檔案：%s"
	MsgMissingFileNotCreated            = "缺少的檔案 %s 未能建立：%s"
	MsgFileIntegrityValid               = "檔案完整性正常：%s"
	MsgSkippingModifiedFile             = "略過已修改的檔案：%s"
	MsgDirectoryCheckedWithModified     = "目錄檢查完成（存在被修改的檔案）：%s"
)

// Update operations messages
const (
	MsgCheckingUpdates            = "正在檢查更新"
	MsgDownloadingInstaller       = "正在下載安裝程式更新"
	MsgInstallerDownloadCompleted = "安裝程式下載完成"
	MsgUpdatingInstaller          = "正在更新安裝程式"
	MsgRemovingInstaller          = "正在移除安裝程式"
	MsgInstallerUpdateCompleted   = "安裝程式更新完成"
	// Past tense: the length, the sha256 and the Ed25519 signature are all checked as the
	// stream arrives, so by the time this is printed the verdict is already in.
	MsgVerifyingBinaryChecksum   = "檢查碼與簽章驗證通過"
	MsgInstallerSavedTo          = "安裝程式已儲存到 %s"
	MsgInstallerDownloadProgress = "已下載 %d%%（%s / %s）"
	// The installer binary outlives the removal on purpose: a process cannot delete the
	// file it is running from, and trying would leave the user with neither.
	MsgInstallerNotRemoved            = "安裝程式本身保留在 %s —— 不再需要時請自行刪除"
	MsgInstallerNotRemovedUnknownPath = "安裝程式本身會保留在原處 —— 不再需要時請自行刪除"
)

// Jaeger storage plugin messages. The plugin is a file mounted into the Jaeger container,
// not an image, so none of the compose output covers it.
const (
	MsgJaegerPluginDownloading       = "正在下載 Jaeger 儲存外掛 %s（%s）"
	MsgJaegerPluginProgress          = "  %s：%d%%（%s / %s）"
	MsgJaegerPluginInstalled         = "Jaeger 儲存外掛已就位：%s"
	MsgJaegerPluginStaleRemoved      = "已移除中斷下載留下的殘留檔案：%s"
	MsgJaegerPluginStaleRemoveFailed = "無法移除殘留檔案 %s：%v"
	MsgJaegerPluginRestarting        = "正在重新啟動 Jaeger 以載入新的儲存外掛"
)

// Post-update verification messages
const (
	MsgVerifyingUpdatedComponents = "正在驗證已更新的 %s 元件"
	MsgComponentUpToDate          = "  %s：與已發布的組建相符"
	MsgComponentNewerThanOffered  = "  %s：比提供更新的版本（登錄檔已前進）"
	MsgComponentMismatch          = "  %s：預期 %s，實際 %s"
	MsgComponentNotVerifiable     = "  %s：沒有已發布的摘要可比較"
	MsgVerificationPassed         = "所有已更新的 %s 元件都與已發布的組建相符"
	MsgVerificationMismatch       = "有 %d 個 %s 元件與已發布的組建不符"
)

// Remove operations messages
const (
	MsgRemovingStack          = "正在移除堆疊：%s"
	MsgStackRemovalCompleted  = "%s 的堆疊移除完成"
	MsgPurgingStack           = "正在清除堆疊：%s"
	MsgStackPurgeCompleted    = "%s 的堆疊清除完成"
	MsgExecutingDockerCompose = "正在執行 docker-compose 指令：%s"
	MsgDockerComposeCompleted = "docker-compose 指令已完成"
	MsgFactoryResetStarting   = "開始出廠重設"
	MsgFactoryResetCompleted  = "出廠重設完成"
	MsgRestoringDefaultEnv    = "正在從內嵌資源還原預設 .env"
	MsgDefaultEnvRestored     = "已還原預設 .env"
)

type Subsystem string

const (
	SubsystemDocker     Subsystem = "docker"
	SubsystemCompose    Subsystem = "compose"
	SubsystemFileSystem Subsystem = "file-system"
	SubsystemUpdate     Subsystem = "update"
)

type SubsystemOperationMessage struct {
	Enter string
	Exit  string
	Error string
}

var SubsystemOperationMessages = map[Subsystem]map[ProcessorOperation]SubsystemOperationMessage{
	SubsystemCompose: {
		ProcessorOperationStart: SubsystemOperationMessage{
			Enter: "正在啟動 %s compose 堆疊",
			Exit:  "已啟動 compose 堆疊 %s",
			Error: "啟動 %s compose 堆疊失敗",
		},
		ProcessorOperationStop: SubsystemOperationMessage{
			Enter: "正在停止 %s compose 堆疊",
			Exit:  "已停止 compose 堆疊 %s",
			Error: "停止 %s compose 堆疊失敗",
		},
		ProcessorOperationRestart: SubsystemOperationMessage{
			Enter: "正在重新啟動 %s compose 堆疊",
			Exit:  "已重新啟動 compose 堆疊 %s",
			Error: "重新啟動 %s compose 堆疊失敗",
		},
		ProcessorOperationUpdate: SubsystemOperationMessage{
			Enter: "正在更新 %s compose 堆疊",
			Exit:  "已更新 compose 堆疊 %s",
			Error: "更新 %s compose 堆疊失敗",
		},
		ProcessorOperationDownload: SubsystemOperationMessage{
			Enter: "正在下載 %s compose 堆疊",
			Exit:  "已下載 compose 堆疊 %s",
			Error: "下載 %s compose 堆疊失敗",
		},
		ProcessorOperationRemove: SubsystemOperationMessage{
			Enter: "正在移除 %s compose 堆疊",
			Exit:  "已移除 compose 堆疊 %s",
			Error: "移除 %s compose 堆疊失敗",
		},
		ProcessorOperationPurge: SubsystemOperationMessage{
			Enter: "正在清除 %s compose 堆疊",
			Exit:  "已清除 compose 堆疊 %s",
			Error: "清除 %s compose 堆疊失敗",
		},
	},
}
