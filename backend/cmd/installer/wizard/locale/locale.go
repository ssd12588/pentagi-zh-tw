package locale

// Common status and UI strings
const (
	// Common status and UI strings
	UIStatistics       = "統計"
	UIStatus           = "狀態："
	UIMode             = "模式："
	UINoConfigSelected = "尚未選擇設定"
	UILoading          = "載入中…"
	UINotImplemented   = "尚未實作"
	UIUnsavedChanges   = "尚未儲存的變更"
	UIConfigSaved      = "設定已儲存"

	// Status labels
	StatusEnabled       = "已啟用"
	StatusDisabled      = "已停用"
	StatusConfigured    = "已設定"
	StatusNotConfigured = "未設定"
	StatusEmbedded      = "內嵌"
	StatusExternal      = "外部"

	// Success/Warning messages
	MessageSearchEnginesNone       = "⚠ 尚未設定搜尋引擎"
	MessageSearchEnginesConfigured = "✓ 已設定 %d 個搜尋引擎"
	MessageDockerConfigured        = "✓ Docker 環境已設定"
	MessageDockerNotConfigured     = "⚠ Docker 環境尚未設定"
)

// Legend constants
const (
	LegendConfigured    = "✓ 已設定"
	LegendNotConfigured = "✗ 未設定"
)

// Common Navigation Actions (always available)
const (
	NavBack       = "Esc：返回"
	NavExit       = "Ctrl+Q：離開"
	NavUpDown     = "↑/↓：捲動／選擇"
	NavLeftRight  = "←/→：移動"
	NavPgUpPgDown = "PgUp/PgDn：翻頁"
	NavHomeEnd    = "Home/End：跳到開頭／結尾"
	NavEnter      = "Enter：繼續"
	NavYn         = "Y/N：接受／拒絕"
	NavCtrlC      = "Ctrl+C：取消"
	NavCtrlS      = "Ctrl+S：儲存"
	NavCtrlR      = "Ctrl+R：重設"
	NavCtrlH      = "Ctrl+H：顯示／隱藏"
	NavTab        = "Tab：自動完成"
	NavSeparator  = " • "
)

// Welcome Screen constants
const (
	// Form interface implementation
	WelcomeFormTitle       = "歡迎使用 PentAGI"
	WelcomeFormDescription = "PentAGI 是一套自主滲透測試平台，運用 AI 技術執行全面的安全評估。"
	WelcomeFormName        = "歡迎"
	WelcomeFormOverview    = `System checks verify:
• Environment configuration file presence
• Docker API accessibility and version compatibility
• Worker environment readiness
• System resources (CPU, memory, disk space)
• Network connectivity for external dependencies

Once all checks pass, proceed through the configuration wizard to set up LLM providers, monitoring, and security tools.

The installer guides you through each component setup with recommendations for different deployment scenarios.`

	// Configuration status messages
	WelcomeConfigurationFailed = "⚠ 檢查未通過：%s"
	WelcomeConfigurationPassed = "✓ 所有系統檢查通過"

	// Workflow steps
	WelcomeWorkflowTitle = "安裝流程："
	WelcomeWorkflowStep1 = "1. 接受使用者授權條款（EULA）"
	WelcomeWorkflowStep2 = "2. 設定 LLM 供應商（OpenAI、Anthropic 等）"
	WelcomeWorkflowStep3 = "3. 設定整合服務（Langfuse、可觀測性）"
	WelcomeWorkflowStep4 = "4. 設定安全性選項"
	WelcomeWorkflowStep5 = "5. 部署並啟動 PentAGI 服務"
	WelcomeSystemReady   = "✓ 系統就緒 —— 按 Enter 繼續"
)

// Troubleshooting on welcome screen constants
const (
	TroubleshootTitle = "系統需求未達成"

	// Environment file issues
	TroubleshootEnvFileTitle = "缺少環境設定檔"
	TroubleshootEnvFileDesc  = "PentAGI 需要 .env 設定檔，但找不到或無法讀取。"
	TroubleshootEnvFileFix   = `To fix:
1. Copy .env.example to .env in your installation directory
2. Edit .env and configure at least one LLM provider API key
3. Ensure the file has read permissions (chmod 644 .env)

Quick fix:
cp .env.example .env && chmod 644 .env`

	// Write permissions
	TroubleshootWritePermTitle = "需要寫入權限"
	TroubleshootWritePermDesc  = "安裝程式需要設定目錄的寫入權限，才能儲存設定並部署服務。"
	TroubleshootWritePermFix   = `To fix:
1. Check directory permissions: ls -la
2. Grant write access: chmod 755 .
3. Or run installer from a writable location
4. Ensure sufficient disk space is available`

	// Docker not installed
	TroubleshootDockerNotInstalledTitle = "未安裝 Docker"
	TroubleshootDockerNotInstalledDesc  = "此系統未安裝 Docker。PentAGI 需要 Docker 才能執行容器。"
	TroubleshootDockerNotInstalledFix   = `To fix:
1. Install Docker Desktop: https://docs.docker.com/get-docker/
2. For Linux: Follow distribution-specific instructions
3. Verify installation: docker --version
4. Ensure docker command is in your PATH`

	// Docker not running
	TroubleshootDockerNotRunningTitle = "Docker 服務未執行"
	TroubleshootDockerNotRunningDesc  = "已安裝 Docker，但服務未執行。Docker 服務必須處於執行狀態。"
	TroubleshootDockerNotRunningFix   = `To fix:
1. Start Docker Desktop (Windows/Mac)
2. Linux: sudo systemctl start docker
3. Check status: docker ps
4. If using DOCKER_HOST, verify the remote daemon is accessible`

	// Docker permission issues
	TroubleshootDockerPermissionTitle = "Docker 權限不足"
	TroubleshootDockerPermissionDesc  = "你的使用者帳號沒有存取 Docker 的權限（Linux 常見）。"
	TroubleshootDockerPermissionFix   = `To fix:
1. Add user to docker group: sudo usermod -aG docker $USER
2. Log out and back in for changes to take effect
3. Or run with sudo (not recommended for production)
4. Verify: docker ps (should work without sudo)`

	// Generic Docker API issues
	TroubleshootDockerAPITitle = "Docker API 連線失敗"
	TroubleshootDockerAPIDesc  = "無法建立與 Docker API 的連線，可能是設定或網路問題。"
	TroubleshootDockerAPIFix   = `To fix:
1. Check DOCKER_HOST environment variable
2. Verify Docker is running: docker version
3. For remote Docker: ensure network connectivity
4. Check firewall settings if using TCP connection
5. Try: export DOCKER_HOST=unix:///var/run/docker.sock`

	// Docker version issues
	TroubleshootDockerVersionTitle = "Docker 版本過舊"
	TroubleshootDockerVersionDesc  = "你的 Docker 版本不相容。PentAGI 需要 Docker 20.0.0 以上。"
	TroubleshootDockerVersionFix   = `To fix:
1. Update Docker to version 20.0.0 or newer
2. Visit https://docs.docker.com/engine/install/

Current version: %s
Required: 20.0.0+`

	// Docker Compose issues
	TroubleshootComposeTitle = "找不到 Docker Compose"
	TroubleshootComposeDesc  = "需要 Docker Compose v2 外掛，但目前無法使用 `docker compose`。"
	TroubleshootComposeFix   = `To fix:
1. Install or update Docker Desktop, or install the Docker Compose v2 plugin for Docker Engine
2. Verify the plugin is available: docker compose version
3. If only legacy docker-compose is installed, install the Docker Compose v2 plugin as well

PentAGI executes "docker compose", so legacy "docker-compose" alone is not sufficient.
Documentation: https://docs.docker.com/compose/install/`

	// Docker Compose version issues
	TroubleshootComposeVersionTitle = "Docker Compose 版本過舊"
	TroubleshootComposeVersionDesc  = "你的 `docker compose` 版本不相容。PentAGI 需要 Docker Compose 1.25.0 以上。"
	TroubleshootComposeVersionFix   = `Current version: %s
Required: 1.25.0+

To fix:
1. Update Docker Desktop or the Docker Compose v2 plugin to a newer version
2. Verify the result with: docker compose version

Documentation: https://docs.docker.com/compose/install/`

	// Worker environment issues
	TroubleshootWorkerTitle = "無法存取 worker 的 Docker 環境"
	TroubleshootWorkerDesc  = "無法連線到 worker 容器所需的 Docker 環境。可能是遠端或本機 Docker 設定問題。"
	TroubleshootWorkerFix   = `To fix:
1. For remote Docker, set env vars before installer:
   export DOCKER_HOST=tcp://remote:2376
   export DOCKER_CERT_PATH=/path/to/certs
   export DOCKER_TLS_VERIFY=1
2. Verify connection: docker -H $DOCKER_HOST ps
3. For local Docker, leave these vars unset
4. Check firewall allows Docker port (2375/2376)
5. Ensure certificates are valid if using TLS`

	// CPU issues
	TroubleshootCPUTitle = "CPU 核心數不足"
	TroubleshootCPUDesc  = "PentAGI 至少需要 2 個 CPU 核心才能正常運作。"
	TroubleshootCPUFix   = `Your system has %d CPU core(s), but 2+ are required.

For virtual machines:
1. Increase CPU allocation in VM settings
2. Ensure host has sufficient resources

Docker Desktop users:
Settings → Resources → CPUs: Set to 2 or more`

	// Memory issues
	TroubleshootMemoryTitle = "記憶體不足"
	TroubleshootMemoryDesc  = "可用記憶體不足以執行所選元件。"
	TroubleshootMemoryFix   = `Memory requirements:
• Base system: 0.5 GB
• PentAGI core: +0.5 GB
• Langfuse (if enabled): +1.5 GB
• Observability (if enabled): +1.5 GB

Total needed: %.1f GB
Available: %.1f GB

To fix:
1. Close unnecessary applications
2. Increase Docker memory limit
3. Disable optional components (Langfuse/Observability)`

	// Disk space issues
	TroubleshootDiskTitle = "磁碟空間不足"
	TroubleshootDiskDesc  = "可用磁碟空間不足以安裝與運作。"
	TroubleshootDiskFix   = `Disk requirements:
• Base installation: 5 GB minimum
• With components: 10 GB + 2 GB per component
• Worker images: 25 GB (includes 6GB+ Kali image)

Required: %.1f GB
Available: %.1f GB

To fix:
1. Free up disk space
2. Use external storage for Docker
3. Prune unused Docker resources:
   docker system prune -a`

	// Network issues
	TroubleshootNetworkTitle = "網路連線失敗"
	TroubleshootNetworkDesc  = "無法連線到必要的外部服務，將無法下載 Docker 映像與更新。"
	TroubleshootNetworkFix   = `Failed checks:
%s

To fix:
1. Verify internet connection: ping docker.io
2. Check DNS resolution: nslookup docker.io
3. If behind proxy, set before running installer:
   export HTTP_PROXY=http://proxy:port
   export HTTPS_PROXY=http://proxy:port
4. For persistent proxy, add to .env:
   PROXY_URL=http://proxy:port
5. Check firewall allows outbound HTTPS (port 443)
6. Try alternative DNS servers if DNS fails`

	// Generic hint at the bottom
	TroubleshootFixHint = "\\n請先解決上述問題，再重新執行安裝程式。"

	// Network failure messages (used in checker/helpers.go)
	NetworkFailureDNS        = "• docker.io 的 DNS 解析失敗"
	NetworkFailureHTTPS      = "• 無法透過 HTTPS 連線到外部服務"
	NetworkFailureDockerPull = "• 無法從登錄檔拉取 Docker 映像"
)

// System Checks constants
const (
	ChecksTitle               = "系統檢查"
	ChecksWarningFailed       = "⚠ 部分檢查未通過"
	CheckEnvironmentFile      = "環境設定檔"
	CheckWritePermissions     = "寫入權限"
	CheckDockerAPI            = "Docker API"
	CheckDockerVersion        = "Docker 版本"
	CheckDockerCompose        = "Docker Compose"
	CheckDockerComposeVersion = "Docker Compose 版本"
	CheckWorkerEnvironment    = "worker 環境"
	CheckSystemResources      = "系統資源"
	CheckNetworkConnectivity  = "網路連線"
)

// EULA Screen constants
const (
	// Form interface implementation
	EULAFormDescription = "PentAGI 使用的法律條款與條件"
	EULAFormName        = "EULA"
	EULAFormOverview    = `Review and accept the End User License Agreement to proceed with PentAGI installation.

The EULA contains:
• Software license terms and usage rights
• Limitation of liability and warranties
• Data collection and privacy policies
• Compliance requirements and restrictions
• Support and maintenance terms

You must scroll through the entire document and accept the terms to continue with the installation process.

Use arrow keys, page up/down, or home/end keys to navigate through the document.`

	// Error and status messages
	EULAErrorLoadingTitle     = "# 載入授權條款失敗\\n\\n無法載入 EULA：%v"
	EULAContentFallback       = "# 授權條款內容\\n\\n%s\\n\\n---\\n\\n*註：Markdown 渲染失敗：%v*"
	EULAConfigurationRead     = "✓ 已檢視使用者授權條款"
	EULAConfigurationAccepted = "✓ 已接受使用者授權條款"
	EULAConfigurationPending  = "⚠ 尚未檢視使用者授權條款"
	EULALoading               = "正在載入授權條款…"
	EULAProgress              = "進度：%d%%"
	EULAProgressComplete      = " • 完成"
)

// Update Overview Screen constants
const (
	UpdateOverviewFormName        = "更新 PentAGI"
	UpdateOverviewFormDescription = "此更新帶來的變更與對應的版本說明"
	UpdateOverviewFormOverview    = `Review what the update will change before applying it.

The overview shows:
• Each stack with the version it is on and the version it moves to
• Every component and whether it will actually change
• The changelog and release notes of every release you cross

Nothing is downloaded or changed on this screen. Continue to apply the update.

Use arrow keys, page up/down, or home/end keys to navigate through the document.`

	UpdateOverviewLoading         = "正在準備更新概覽…"
	UpdateOverviewConfigurationOK = "✓ 有更新可用"
	UpdateOverviewConfigurationNo = "✓ 所有元件皆為最新"

	// Document
	UpdateOverviewHeading         = "# 更新概覽"
	UpdateOverviewNothingToDo     = "一切皆為最新，沒有需要套用的變更。"
	UpdateOverviewCheckFailed     = "上次更新檢查未完成，因此這裡沒有可顯示的內容。\\n\\n套用更新時，會拉取 compose 檔目前指向的版本。"
	UpdateOverviewStackHeading    = "## %s"
	UpdateOverviewVersionMove     = "**%s → %s**"
	UpdateOverviewVersionTarget   = "**→ %s**"
	UpdateOverviewVersionUnknown  = "*版本未知 —— 此安裝無法對應到任何發行版本。*"
	UpdateOverviewVersionMixed    = "*各元件來自不同版本，因此上方顯示的是其中最舊的版本。*"
	UpdateOverviewComponentsTitle = "元件："
	UpdateOverviewComponentChange = "- **%s** (%s/%s) — %s"
	UpdateOverviewWillChange      = "將被更新"
	UpdateOverviewNoChange        = "unchanged"
	UpdateOverviewCannotVerify    = "無法驗證 —— 沒有可比較的對象"
	UpdateOverviewReleaseHeading  = "### %s"
	UpdateOverviewReleasePreview  = "### %s（預覽）"
	UpdateOverviewReleaseDate     = "*發行於 %s*"
	UpdateOverviewReleasesCut     = "*已省略較舊版本 —— 僅顯示最近 %d 個。*"
	UpdateOverviewUpToDateTitle   = "## 已是最新版本"
	UpdateOverviewRenderFallback  = "%s\\n\\n---\\n\\n*註：Markdown 渲染失敗：%v*"
)

// Installer Update Screen constants.
//
// The wording is deliberate and was wrong before: the screen this replaced promised to
// "replace the current installer" and "exit for manual restart", and no version of this
// installer has ever done either. What happens is a download of one verified file.
const (
	InstallerUpdateFormOverview = `Download the installer build the update server offers for this machine.

The file is written next to your installation under its own versioned name, so the installer you are running now is left untouched. Its length, sha256 and signature are checked as it arrives; anything that fails is deleted rather than kept.

Replacing the installer you use is your decision to make afterwards — run the downloaded file when you are ready.`

	InstallerUpdateHelpTitle       = "更新安裝程式"
	InstallerUpdateConfigurationOK = "✓ 有較新的版本可用"
	InstallerUpdateConfigurationNo = "✓ 安裝程式已是最新"

	// Left panel: what is on offer, before anything is fetched.
	InstallerUpdateAsking         = "正在向更新伺服器查詢可用的版本…"
	InstallerUpdateLoadFailed     = "無法取得套件說明：%v"
	InstallerUpdateCurrentVersion = "• 目前執行版本：%s"
	InstallerUpdateOfferedVersion = "• 可更新版本：%s"
	InstallerUpdatePlatform       = "• 平台：%s/%s"
	InstallerUpdateSize           = "• 下載大小：%s"
	InstallerUpdateTarget         = "• 將儲存為：%s"
	InstallerUpdateManualNote     = "正在執行的安裝程式不會被取代。下載完成時會顯示指令，請自行把下載的檔案移到定位。"
	InstallerUpdateSameVersion    = "提供的版本就是目前執行中的版本，下載不會有任何改變。"
	InstallerUpdatePressEnter     = "按 Enter 開始下載"

	// The only thing worth a confirmation: a file under that exact name is already here,
	// and it may be one somebody put there on purpose. Downloading over nothing needs no
	// permission and is not asked about.
	InstallerUpdateAlreadyHere     = "此名稱的檔案已存在 —— 可能是先前的下載，或下載中斷留下的檔案。"
	InstallerUpdateOverwritePrompt = "要重新下載並取代 %s 嗎？"
	InstallerUpdatePressYN         = "按 y 下載並取代，按 n 保留磁碟上既有的檔案"
	InstallerUpdateKeptExisting    = "已保留磁碟上既有的檔案，未進行下載。"

	// Left panel: the download itself. There is no "completed" here — the operation
	// reports where the file landed and that it verified, and a second line saying the
	// same thing reads as a second thing having happened.
	InstallerUpdateInProgress = "正在下載安裝程式版本…\\n"
	InstallerUpdateFailed     = "下載安裝程式版本失敗"

	// What to do with the file afterwards. The installer never does this itself.
	InstallerUpdateFileIsAt    = "組建位於 %s"
	InstallerUpdateMoveTitle   = "要開始使用，請執行："
	InstallerUpdateMoveUnknown = "準備好時，請將它覆蓋你啟動的安裝程式 —— 安裝程式不會自行取代自己。"
	InstallerUpdateWindowsNote = "Windows 不會取代正在執行的程式：請先關閉此安裝程式，再執行該指令。"
	InstallerUpdateMoveIndent  = "    %s"
)

// Main Menu Screen constants
const (
	MainMenuTitle       = "PentAGI 設定"
	MainMenuDescription = "設定 PentAGI 的所有元件與選項"
	MainMenuName        = "主選單"
	MainMenuOverview    = `Welcome to PentAGI Configuration Center.

Configure essential components:
• LLM Providers - AI language models for autonomous testing
• Monitoring - Observability and analytics platforms
• Tools - Additional capabilities for enhanced testing
• System Settings - Environment and deployment options

Navigate through each section to complete your PentAGI setup.`

	MenuTitle        = "設定選單"
	MenuSystemStatus = "系統狀態"
)

// Main Menu Status Labels (not used)
const (
	MainMenuStatusPentagiRunning     = "PentAGI 已在執行中"
	MainMenuStatusPentagiNotRunning  = "準備啟動 PentAGI 服務"
	MainMenuStatusUpToDate           = "PentAGI 已是最新版本"
	MainMenuStatusUpdatesAvailable   = "有可用更新"
	MainMenuStatusReadyToStart       = "準備開始"
	MainMenuStatusAllServicesRunning = "所有服務均在執行中"
	MainMenuStatusNoUpdatesAvailable = "沒有可用更新"
)

// LLM Providers Screen constants
const (
	LLMProvidersTitle       = "LLM 供應商設定"
	LLMProvidersDescription = "設定 AI 代理使用的大型語言模型供應商"
	LLMProvidersName        = "LLM 供應商"
	LLMProvidersOverview    = `PentAGI uses specialized AI agents (researcher, developer, executor, pentester) that require different LLM capabilities for optimal penetration testing results.

Why multiple providers matter:
• Agent Specialization: Different agents benefit from models optimized for reasoning, coding, or analysis
• Cost Efficiency: Mix expensive reasoning models (o3, grok-4, claude-sonnet-4, gemini-2.5-pro) for complex tasks with cheaper models for simple operations
• Performance Optimization: Each provider excels in different areas - OpenAI for medium tasks, Anthropic for complex tasks, Gemini for saving costs

Provider Selection Guide:
• Cloud Production: OpenAI + Anthropic + Gemini for industry-leading performance and reliability
• Enterprise/Compliance: AWS Bedrock for SOC2, HIPAA, and access to multiple model families
• Privacy/On-premises: Ollama or vLLM with Llama 3.1, Qwen3, or other open models for complete data control

Ready-to-use configurations for OpenRouter, DeepInfra, vLLM, Ollama, and other providers are available in the /opt/pentagi/conf/ directory inside the container`
)

// LLM Provider titles and descriptions
const (
	LLMProviderOpenAI        = "OpenAI"
	LLMProviderAnthropic     = "Anthropic"
	LLMProviderGemini        = "Google Gemini"
	LLMProviderBedrock       = "AWS Bedrock"
	LLMProviderOllama        = "Ollama"
	LLMProviderDeepSeek      = "DeepSeek"
	LLMProviderGLM           = "GLM（智譜 AI）"
	LLMProviderKimi          = "Kimi（Moonshot AI）"
	LLMProviderQwen          = "Qwen（阿里雲）"
	LLMProviderMiniMax       = "MiniMax"
	LLMProviderMistral       = "Mistral AI"
	LLMProviderXAI           = "xAI Grok"
	LLMProviderCustom        = "自訂"
	LLMProviderOpenAIDesc    = "業界領先的 GPT 模型，整體表現優異"
	LLMProviderAnthropicDesc = "具備優異推理與安全特性的 Claude 模型"
	LLMProviderGeminiDesc    = "Google 具備廣泛知識的進階多模態模型"
	LLMProviderBedrockDesc   = "企業級 AWS 存取，可使用多個基礎模型供應商"
	LLMProviderOllamaDesc    = "本機與雲端開源模型，兼顧隱私與彈性"
	LLMProviderDeepSeekDesc  = "具備強大推理與多語言能力的中文 AI 模型"
	LLMProviderGLMDesc       = "智譜 AI 的 GLM 模型，適合中英文任務"
	LLMProviderKimiDesc      = "Moonshot AI 長上下文模型，適合文件分析"
	LLMProviderQwenDesc      = "阿里雲 Qwen 系列模型，適合多語言任務"
	LLMProviderMiniMaxDesc   = "MiniMax M 系列模型，適合代理推理與長上下文任務"
	LLMProviderMistralDesc   = "透過 Mistral 對話 API 使用 Medium、Small 與 Ministral 模型"
	LLMProviderXAIDesc       = "透過 xAI 對話 API 使用 Grok 模型"
	LLMProviderCustomDesc    = "自訂 OpenAI 相容端點，提供最大彈性"
)

// Provider-specific help text
const (
	LLMFormOpenAIHelp = `OpenAI delivers industry-leading models with cutting-edge reasoning capabilities perfect for sophisticated penetration testing.

Default PentAGI Models:
• o1, o4-mini: Advanced reasoning models for complex vulnerability analysis and strategic planning
• GPT-4.1, GPT-4.1-mini: Flagship models optimized for exploit development and code generation
• Automatic model selection based on agent type and task complexity

Key Advantages:
• Most advanced reasoning capabilities with step-by-step analysis (o-series models)
• Excellent coding abilities for custom exploit development and payload generation
• Reliable performance with consistent uptime and extensive API documentation
• Proven track record in security research and penetration testing scenarios

Best for: Production environments requiring cutting-edge AI capabilities, teams prioritizing performance over cost
Cost: Premium pricing, but optimized configurations balance cost with quality

Setup: Get your API key from https://platform.openai.com/api-keys`

	LLMFormAnthropicHelp = `Anthropic Claude models excel in safety-conscious penetration testing with superior reasoning and analytical capabilities.

Default PentAGI Models:
• Claude Sonnet-4: Premium reasoning model for complex security analysis and strategic vulnerability assessment
• Claude 3.5 Haiku: High-speed model optimized for rapid information gathering and simple parsing tasks
• Balanced cost-performance ratio across all security testing scenarios

Key Advantages:
• Exceptional safety and ethics focus - reduces harmful output while maintaining security testing effectiveness
• Superior reasoning for methodical vulnerability analysis and systematic penetration testing approaches
• Large context windows ideal for analyzing extensive codebases and configuration files
• Excellent at understanding complex security contexts and regulatory compliance requirements

Best for: Security teams prioritizing responsible testing practices, compliance-focused environments, detailed analysis
Cost: Mid-range pricing with excellent value for reasoning-heavy security workflows

Setup: Get your API key from https://console.anthropic.com/`

	LLMFormAnthropicFederatedHelp = `Federated (Enterprise) authentication uses a workload identity token instead of a static API key.

PentAGI presents a short-lived identity token (inline via the Identity Token field, or read at runtime from the Identity Token File) together with the Organization, Workspace, Service Account and Federation Rule that map it to Anthropic access.

Use this when your organization issues federated credentials (for example from a cloud workload identity provider) rather than a long-lived API key. All four IDs plus one of the two token fields are required for federated auth to be considered configured.

Switch back to "API 金鑰" for the standard console API key.`

	LLMFormGeminiHelp = `Google Gemini combines multimodal capabilities with advanced reasoning, perfect for comprehensive security assessments.

Default PentAGI Models:
• gemini-3.5-flash-lite: planning, the primary agent loop, code, and other tool-using agents
• gemini-3.1-flash-lite: simple output, JSON, reflection, search and enrichment

Key Advantages:
• Multimodal support enables analysis of screenshots, network diagrams, and security documentation
• Competitive pricing with generous rate limits for development and testing environments
• Large context windows (up to 2M tokens) for analyzing massive codebases and system configurations
• Strong performance in code analysis and vulnerability identification across multiple programming languages

Best for: Budget-conscious teams, development environments, scenarios requiring image/document analysis
Cost: Most cost-effective option among major cloud providers with excellent performance/price ratio

Setup: Get your API key from https://aistudio.google.com/app/apikey`

	LLMFormBedrockHelp = `AWS Bedrock provides enterprise-grade access to 20+ foundation models with multiple authentication methods and enhanced security.

Default PentAGI Models:
• Claude Sonnet-4.5 (via Bedrock): Premium reasoning model with AWS enterprise security and extended thinking capabilities
• OpenAI GPT OSS 120B: Strong reasoning model for scientific analysis and complex security tasks
• Claude Haiku-4.5, DeepSeek V3.2, Qwen3-32B: Efficient models for specific agent roles and cost optimization
• Access to Amazon Nova (multimodal), Mistral, Moonshot, and more through single unified interface

Authentication Methods (priority order):
1. Default AWS Auth (BEDROCK_DEFAULT_AUTH=true): Use AWS SDK credential chain - recommended for EC2/ECS/Lambda
2. Bearer Token (BEDROCK_BEARER_TOKEN): Token-based authentication for custom auth scenarios
3. Static Credentials (ACCESS_KEY + SECRET_KEY): Traditional IAM credentials for development and testing

Key Advantages:
• Enterprise compliance: SOC2, HIPAA, FedRAMP certifications with data residency and governance controls
• Multi-provider access: 20+ models from Anthropic, Amazon, OpenAI, Qwen, DeepSeek, Cohere, Mistral, Moonshot
• Flexible authentication: Three methods to suit different deployment scenarios and security requirements
• Enhanced security: VPC integration, CloudTrail logging, IAM controls, private endpoints for complete isolation
• Regional deployment: Deploy in preferred AWS regions for latency optimization and data sovereignty

Best for: Enterprise environments, regulated industries, teams requiring compliance controls and flexible authentication
Cost: Competitive pricing with provisioned throughput options, but new accounts have restrictive rate limits (2-20 req/min)
Important: Request quota increases through AWS Service Quotas console for production penetration testing workflows

Setup: Choose authentication method and configure credentials. Verify rate limits at https://docs.aws.amazon.com/bedrock/`

	LLMFormOllamaHelp = `Ollama supports two deployment scenarios for complete flexibility.

Scenario 1: Local Ollama Server (Self-Hosted)
• Run Ollama on your own hardware (8GB+ RAM recommended, GPU optional but beneficial)
• Complete data privacy - all processing happens locally
• Zero ongoing costs - only infrastructure
• No API key needed - authentication handled by network access
• Setup: Install from https://ollama.ai/ and configure OLLAMA_SERVER_URL=http://ollama-server:11434

Scenario 2: Ollama Cloud (Managed Service)
• Cloud-hosted models without local infrastructure requirements
• No hardware needed - models run on Ollama's infrastructure
• Pay-per-use pricing with free tier available
• API key required - generate at https://ollama.com/settings/keys
• Setup: Register at https://ollama.com, configure OLLAMA_SERVER_URL=https://ollama.com + OLLAMA_SERVER_API_KEY=your_key

Default PentAGI Models:
• Llama 3.1:8b, Qwen3:32b, and other open models
• Customizable - switch between 100+ available models
• Model auto-download and loading options for convenience

Key Advantages:
• Dual deployment options: Choose between privacy (local) and convenience (cloud)
• Cost flexibility: Zero ongoing costs for local, pay-per-use for cloud
• Extensive model library: Access to latest open-source models (Llama, Qwen, Mistral, Gemma, and more)
• Air-gapped support: Local deployment works in isolated networks

Best for: Privacy-focused teams (local), budget-conscious deployments (cloud), organizations with data sovereignty requirements
Setup options: Local installation from https://10.10.10.10:11434 or cloud registration at https://ollama.com`

	LLMFormDeepSeekHelp = `DeepSeek provides advanced AI models with strong reasoning capabilities and multilingual support.

Default PentAGI Models:
• deepseek-flash: Cost-efficient general-purpose model for dialogue, code generation, and tool calling
• deepseek-v4-pro: Higher-tier reasoning model for complex logic, mathematical reasoning, and security analysis
• Cost-effective pricing with competitive performance compared to leading models

Key Advantages:
• Strong coding and reasoning capabilities for security analysis and exploit development
• Multilingual support (Chinese and English) for international penetration testing scenarios
• Competitive pricing with excellent performance-to-cost ratio
• OpenAI-compatible API for seamless integration

LiteLLM Integration:
• Set Provider Name to 'deepseek' when using LiteLLM proxy
• Enables model prefix (e.g., deepseek/deepseek-flash) without modifying config.yml
• Optional for direct DeepSeek API usage

Best for: Teams requiring multilingual support, cost-conscious deployments, Chinese language security testing
Cost: Highly competitive pricing with strong performance characteristics

Setup: Get your API key from https://platform.deepseek.com/`

	LLMFormGLMHelp = `GLM from Zhipu AI provides advanced language models with strong NLP and reasoning capabilities developed by Tsinghua University.

Default PentAGI Models:
• GLM-4-Air: High performance general dialogue model optimized for regular tasks and tool calling
• GLM-4-Plus: Flagship model with strong reasoning and code generation capabilities
• GLM-Z1-Plus: Advanced reasoning model with deep analysis capabilities for security research

Key Advantages:
• Exceptional Chinese and English NLP capabilities
• Strong performance in multilingual security testing and analysis scenarios
• GLM-4 and GLM-Z1 model families with enhanced reasoning and coding
• OpenAI-compatible API for easy integration

Alternative API Endpoints:
• International: https://api.z.ai/api/paas/v4 (default)
• China: https://open.bigmodel.cn/api/paas/v4
• Coding-specific: https://api.z.ai/api/coding/paas/v4

LiteLLM Integration:
• Set Provider Name to 'zai' when using LiteLLM proxy
• Enables model prefix (e.g., zai/glm-4) without modifying config.yml
• Optional for direct GLM API usage

Best for: Chinese and English multilingual penetration testing, teams operating in Asian markets
Cost: Competitive pricing with good performance for multilingual tasks

Setup: Get your API key from https://open.bigmodel.cn/`

	LLMFormKimiHelp = `Kimi from Moonshot AI provides ultra-long context models perfect for analyzing extensive codebases and documentation.

Default PentAGI Models:
• Moonshot-v1-8k: Long-context model supporting up to 8K tokens for general dialogue
• Kimi-k2.5: Advanced model with strong reasoning and document understanding
• Optimized for processing large volumes of text and code

Key Advantages:
• Ultra-long context windows (up to 1M tokens) for comprehensive codebase analysis
• Strong Chinese and English language support for multilingual penetration testing
• Cost-effective for document-heavy security assessments and threat intelligence analysis
• Excellent at understanding complex system architectures and long-form technical documentation

Alternative API Endpoints:
• International: https://api.moonshot.ai/v1 (default)
• China: https://api.moonshot.cn/v1

LiteLLM Integration:
• Set Provider Name to 'moonshot' when using LiteLLM proxy
• Enables model prefix (e.g., moonshot/kimi-k2.6) without modifying config.yml
• Optional for direct Kimi API usage

Best for: Large codebase analysis, document-heavy assessments, teams needing extended context for security research
Cost: Competitive pricing with excellent value for long-context use cases

Setup: Get your API key from https://platform.moonshot.ai/`

	LLMFormQwenHelp = `Qwen from Alibaba Cloud Model Studio (DashScope) provides powerful multilingual models with multimodal capabilities.

Default PentAGI Models:
• Qwen-Turbo: Fastest lightweight model for high-frequency tasks and real-time response scenarios
• Qwen-Plus: Balanced performance model for general dialogue, code generation, and tool calling
• Qwen-Max: Flagship reasoning model with strong instruction following and complex task handling
• QwQ-Plus: Deep reasoning model with extended chain-of-thought for complex logic analysis

Key Advantages:
• Strong multilingual support (Chinese, English, and multiple other languages)
• Multimodal capabilities with Qwen-VL for visual security analysis
• Alibaba Cloud integration for enterprise deployments
• DashScope ecosystem with additional AI services and tools
• Qwen2.5, Qwen3, and QwQ model families with various sizes and specializations

Alternative API Endpoints:
• US: https://dashscope-us.aliyuncs.com/compatible-mode/v1 (default)
• Singapore: https://dashscope-intl.aliyuncs.com/compatible-mode/v1
• China: https://dashscope.aliyuncs.com/compatible-mode/v1

LiteLLM Integration:
• Set Provider Name to 'dashscope' when using LiteLLM proxy
• Enables model prefix (e.g., dashscope/qwen-plus) without modifying config.yml
• Optional for direct Qwen API usage

Best for: Teams operating in Asian markets, multilingual security testing, visual analysis with Qwen-VL, Alibaba Cloud ecosystem integration
Cost: Competitive pricing with flexible tiers for different use cases

Setup: Get your API key from https://dashscope.console.aliyun.com/`

	LLMFormMiniMaxHelp = `MiniMax provides the M-series of agentic models with very large context windows, accessible through an OpenAI-compatible API.

Default PentAGI Models:
• MiniMax-M3: Latest flagship model (~1M token context) for agentic reasoning, tool use, coding, and long-context tasks
• MiniMax-M2.7: Previous-generation model with strong reasoning and coding capabilities
• MiniMax-M2.7-highspeed: Low-latency variant of M2.7 for fast response scenarios

Key Advantages:
• Very large context window on M3 (~1M tokens) for long codebases and reports
• Strong agentic tool-calling, JSON output, and streaming support
• OpenAI- and Anthropic-compatible endpoints for drop-in integration

API Endpoint:
• Global: https://api.minimax.io/v1 (default)

LiteLLM Integration:
• Set Provider Name to 'minimax' when using a LiteLLM proxy
• Enables model prefix (e.g., minimax/MiniMax-M3) without modifying config.yml
• Optional for direct MiniMax API usage

Best for: Agentic security workflows that benefit from very large context and reliable tool calling
Cost: Competitive per-token pricing across the M-series

Setup: Get your API key from https://platform.minimax.io/`

	LLMFormMistralHelp = `Mistral AI serves its models through an OpenAI-compatible chat API.

Default PentAGI Models:
• mistral-medium-latest: planning and advice
• mistral-large-latest: primary and execution agents
• mistral-small-latest: utility and read-heavy agents

API Endpoint:
• https://api.mistral.ai/v1 (default)

LiteLLM Integration:
• Set Provider Name to 'mistral' when using a LiteLLM proxy
• Enables model prefix (e.g., mistral/mistral-small-latest) without modifying config.yml
• Optional for direct Mistral API usage

Setup: Get your API key from https://console.mistral.ai/`

	LLMFormXAIHelp = `xAI serves its Grok models through an OpenAI-compatible chat API.

Default PentAGI Models:
• grok-4.3: planning, the primary agent loop, code, and other tool-using agents
• grok-4.20-0309-non-reasoning: simple output, JSON, reflection, and enrichment

API Endpoint:
• https://api.x.ai/v1 (default)

LiteLLM Integration:
• Set Provider Name to 'xai' when using a LiteLLM proxy
• Enables model prefix (e.g., xai/grok-4.3) without modifying config.yml
• Optional for direct xAI API usage

Setup: Get your API key from https://console.x.ai/`

	LLMFormCustomHelp = `Configure any OpenAI-compatible API endpoint for maximum flexibility and integration with existing infrastructure.

Ready-to-use Configurations:
• vLLM deployments: High-throughput on-premises inference with optimal GPU utilization
• OpenRouter: Access 200+ models from multiple providers through single API with competitive pricing
• DeepInfra: Serverless inference for popular open models with pay-per-use pricing
• Together AI, Groq, Fireworks: Alternative cloud providers with specialized performance optimizations
• LiteLLM Proxy: Universal gateway to 100+ providers with load balancing and unified interface (use LLM_SERVER_PROVIDER for model prefixing)
• Some reasoning models and LLM providers may require preserving reasoning content while using tool calls (LLM_SERVER_PRESERVE_REASONING=true)

Popular On-Premises Options:
• vLLM: Production-grade serving for Qwen, Llama, Mistral models with batching and GPU optimization
• LocalAI: OpenAI-compatible API wrapper for various local models and embedding services
• Text Generation WebUI: Community-favorite interface with extensive model support and fine-tuning capabilities
• Hugging Face TGI: Enterprise text generation inference with auto-scaling and monitoring

Key Advantages:
• Unlimited flexibility: Use any OpenAI-compatible endpoint or service
• Cost optimization: Choose providers with competitive pricing or deploy models on your own infrastructure
• Vendor independence: Avoid lock-in with ability to switch between providers and models seamlessly
• Custom fine-tuning: Deploy specialized models trained on your security testing scenarios

Best for: Teams with specific model requirements, cost optimization needs, or existing LLM infrastructure
LiteLLM Integration: Set LLM_SERVER_PROVIDER to match your provider name (e.g., "openrouter", "moonshot") to use the same config files with both direct API access and LiteLLM proxy
Examples available: Pre-configured setups for major providers in /opt/pentagi/conf/ directory inside the container`
)

// LLM Provider Form field labels and descriptions
const (
	LLMFormFieldBaseURL           = "基底 URL"
	LLMFormFieldAPIKey            = "API 金鑰"
	LLMFormFieldDefaultAuth       = "使用預設 AWS 驗證"
	LLMFormFieldBearerToken       = "Bearer 權杖"
	LLMFormFieldAccessKey         = "存取金鑰 ID"
	LLMFormFieldSecretKey         = "私密存取金鑰"
	LLMFormFieldSessionToken      = "工作階段權杖"
	LLMFormFieldRegion            = "區域"
	LLMFormFieldModel             = "模型"
	LLMFormFieldConfigPath        = "設定檔路徑"
	LLMFormFieldAPIType           = "API 類型"
	LLMFormFieldAPIVersion        = "API 版本"
	LLMFormFieldPreserveReasoning = "保留推理內容"
	LLMFormFieldProviderName      = "供應商名稱"
	LLMFormFieldPullTimeout       = "模型拉取逾時"
	LLMFormFieldPullEnabled       = "自動拉取模型"
	LLMFormFieldLoadModelsEnabled = "從伺服器載入模型"
	LLMFormBaseURLDesc            = "此供應商的 API 端點 URL"
	LLMFormAPIKeyDesc             = "用於驗證的 API 金鑰"
	LLMFormDefaultAuthDesc        = "使用 AWS SDK 預設憑證鏈（環境變數、EC2 角色、~/.aws/credentials）—— 優先度最高"
	LLMFormBearerTokenDesc        = "用於驗證的 Bearer 權杖 —— 優先於靜態憑證"
	LLMFormAccessKeyDesc          = "以靜態憑證驗證時使用的 AWS 存取金鑰 ID"
	LLMFormSecretKeyDesc          = "以靜態憑證驗證時使用的 AWS 私密存取金鑰"
	LLMFormSessionTokenDesc       = "臨時憑證的 AWS 工作階段權杖（選填，需搭配靜態憑證使用）"
	LLMFormRegionDesc             = "Bedrock 服務使用的 AWS 區域"
	LLMFormModelDesc              = "此供應商使用的預設模型"
	LLMFormConfigPathDesc         = "此主機上的設定檔，或從建議清單中選擇容器內的設定檔（選填）"
	LLMFormAPITypeDesc            = "位址慣例：Azure OpenAI 部署請填 azure 或 azure_ad；一般 OpenAI 相容端點請留空"
	LLMFormAPIVersionDesc         = "Azure 部署所需的 api-version；一般端點會忽略此值"
	LLMFormPreserveReasoningDesc  = "在多輪對話中保留推理內容（部分供應商要求）"
	LLMFormProviderNameDesc       = "模型名稱的供應商前綴（適用於 LiteLLM 代理）"
	LLMFormPullTimeoutDesc        = "下載模型的逾時秒數（預設 600）"
	LLMFormPullEnabledDesc        = "啟動時自動下載所需模型"
	LLMFormLoadModelsEnabledDesc  = "從 Ollama 伺服器載入可用模型清單"
	LLMFormOllamaAPIKeyDesc       = "Ollama 雲端 API 金鑰（選填；使用本機 Ollama 可留空）"

	// Anthropic authentication mode (API key vs federated / enterprise)
	LLMAnthropicAuthModeTitle       = "驗證方式"
	LLMAnthropicAuthModeDesc        = "選擇 PentAGI 對 Anthropic 的驗證方式"
	LLMAnthropicAuthModeAPIKey      = "API 金鑰"
	LLMAnthropicAuthModeFederated   = "聯邦驗證（企業版）"
	LLMFormFieldAnthropicOrgID      = "組織 ID"
	LLMFormAnthropicOrgIDDesc       = "請求所屬的 Anthropic 組織"
	LLMFormFieldAnthropicWorkspace  = "工作區 ID"
	LLMFormAnthropicWorkspaceDesc   = "用量與限額所屬的 Anthropic 工作區"
	LLMFormFieldAnthropicServiceAcc = "服務帳號 ID"
	LLMFormAnthropicServiceAccDesc  = "聯邦身分驗證所代表的服務帳號"
	LLMFormFieldAnthropicIDToken    = "身分權杖"
	LLMFormAnthropicIDTokenDesc     = "聯邦身分權杖（此項與權杖檔案擇一使用）"
	LLMFormFieldAnthropicIDTokenF   = "身分權杖檔案"
	LLMFormAnthropicIDTokenFDesc    = "存放聯邦身分權杖的檔案路徑（執行時讀取）"
	LLMFormFieldAnthropicFedRule    = "聯邦規則 ID"
	LLMFormAnthropicFedRuleDesc     = "將身分權杖對應到 Anthropic 存取權的聯邦規則"
)

// LLM Provider Form status messages
const (
	LLMProviderFormTitle       = "LLM 供應商 %s 設定"
	LLMProviderFormDescription = "設定你的大型語言模型供應商選項"
	LLMProviderFormName        = "LLM 供應商 %s"
	LLMProviderFormOverview    = `Agent Role Assignment:
• Primary Agent & Pentester: Use reasoning models (o3, grok-4, claude-sonnet-4, gemini-2.5-pro) for complex vulnerability analysis
• Assistant & Adviser: Advanced models (o4-mini, claude-sonnet-4) for strategic planning and recommendations
• Coder & Installer: Precision models (gpt-4.1, claude-sonnet-4) for exploit development and system configuration
• Searcher & Enricher: Fast models (gpt-4.1-mini, claude-3.5-haiku, gemini-2.0-flash-lite) for information gathering
• Simple tasks: Lightweight models for JSON parsing and basic operations

Performance Considerations:
• Reasoning models provide step-by-step analysis but are slower and more expensive
• Standard models offer faster responses suitable for high-frequency agent interactions
• Each agent type uses provider-specific model configurations optimized for security testing workflows

Your configuration will determine which models each agent uses for different penetration testing scenarios.`
)

// Monitoring Screen
const (
	MonitoringTitle       = "監控設定"
	MonitoringDescription = "設定監控與可觀測性平台，取得完整的系統洞察"
	MonitoringName        = "監控"
	MonitoringOverview    = `Comprehensive monitoring and observability for production-ready deployments.

Why monitoring matters:
• Track performance bottlenecks: Identify slow LLM calls, database queries, and system resources
• Debug issues faster: Detailed traces help diagnose problems across distributed components
• Optimize costs: Monitor token usage patterns and optimize expensive LLM interactions
• Production readiness: Essential for reliable operation in critical environments

Platform Options:
Langfuse: Specialized LLM observability with conversation tracking, prompt engineering insights, and cost analytics
Observability: Full-stack monitoring with metrics, traces, logs, and alerting for infrastructure and application health

Quick Setup:
• Development: Enable Langfuse for LLM insights only
• Production: Enable both platforms for comprehensive monitoring
• Cost-conscious: Use embedded modes to avoid external service fees`
)

// Langfuse Integration constants
const (
	MonitoringLangfuseFormTitle       = "Langfuse 設定"
	MonitoringLangfuseFormDescription = "用於 LLM 監控的 Langfuse 整合設定"
	MonitoringLangfuseFormName        = "Langfuse"
	MonitoringLangfuseFormOverview    = `Langfuse provides:
• Complete conversation tracking
• Model performance metrics
• Cost monitoring and optimization
• User behavior analytics
• Debug traces for AI interactions

Choose between embedded instance or external connection.`

	// Deployment types
	MonitoringLangfuseEmbedded = "內嵌伺服器"
	MonitoringLangfuseExternal = "外部伺服器"
	MonitoringLangfuseDisabled = "已停用"

	// Form fields
	MonitoringLangfuseDeploymentType     = "部署類型"
	MonitoringLangfuseDeploymentTypeDesc = "選擇 Langfuse 的部署方式"
	MonitoringLangfuseBaseURL            = "伺服器 URL"
	MonitoringLangfuseBaseURLDesc        = "Langfuse 伺服器位址（例如 https://cloud.langfuse.com）"
	MonitoringLangfuseProjectID          = "專案 ID"
	MonitoringLangfuseProjectIDDesc      = "Langfuse 中的專案識別碼"
	MonitoringLangfusePublicKey          = "公開金鑰"
	MonitoringLangfusePublicKeyDesc      = "存取專案用的公開 API 金鑰"
	MonitoringLangfuseSecretKey          = "私密金鑰"
	MonitoringLangfuseSecretKeyDesc      = "存取專案用的私密 API 金鑰"
	MonitoringLangfuseListenIP           = "監聽 IP"
	MonitoringLangfuseListenIPDesc       = "Docker 連接埠對應使用的綁定位址（例如 0.0.0.0 表示對所有網路介面開放）"
	MonitoringLangfuseListenPort         = "監聽連接埠"
	MonitoringLangfuseListenPortDesc     = "Langfuse Web 介面對外暴露的 TCP 連接埠"

	// Admin settings for embedded
	MonitoringLangfuseAdminEmail        = "管理員電子郵件"
	MonitoringLangfuseAdminEmailDesc    = "登入 Langfuse 管理介面所使用的電子郵件"
	MonitoringLangfuseAdminPassword     = "管理員密碼"
	MonitoringLangfuseAdminPasswordDesc = "登入 Langfuse 管理介面的密碼"
	MonitoringLangfuseAdminName         = "管理員使用者名稱"
	MonitoringLangfuseAdminNameDesc     = "Langfuse 中的管理員使用者名稱"
	MonitoringLangfuseLicenseKey        = "企業授權金鑰"
	MonitoringLangfuseLicenseKeyDesc    = "Langfuse 企業授權金鑰（選填）"

	// Help text
	MonitoringLangfuseModeGuide    = "選擇部署方式：內嵌（本機控管）、外部（雲端或既有服務）、停用（不使用分析）"
	MonitoringLangfuseEmbeddedHelp = `Embedded deploys complete Langfuse stack:
• PostgreSQL + ClickHouse databases
• MinIO S3 storage + Redis cache
• Full LLM conversation tracking
• Cost analysis and performance metrics
• Private data stays on your server

Resource requirements:
• ~2GB RAM, 5GB disk space minimum
• Additional storage for conversation logs
• Automatic setup and maintenance

Best for: Teams wanting data privacy, custom configurations, or no external dependencies. All analytics data stored locally with full administrative control.

Default admin access:
• Web UI: http://localhost:4000
• Login: admin@pentagi.com
• Password: password (change required)`
	MonitoringLangfuseExternalHelp = `External connects to cloud.langfuse.com or your existing Langfuse server:

• No local infrastructure needed
• Managed updates and maintenance
• Shared analytics across teams
• Enterprise features available
• Data stored on external provider

Setup requirements:
• Langfuse account and API keys
• Internet connectivity required
• Project ID and authentication keys

Best for: Teams using cloud services, wanting managed infrastructure, or integrating with existing Langfuse deployments across organizations.`
	MonitoringLangfuseDisabledHelp = `Langfuse is disabled. Without LLM observability you will not have:

• Conversation history tracking
• Token usage and cost analysis
• Model performance metrics
• Debug traces for AI interactions
• User behavior analytics
• Prompt engineering insights

Consider enabling for production use
to monitor AI agent performance and
optimize costs effectively.`
)

// Graphiti Integration constants
const (
	MonitoringGraphitiFormTitle       = "Graphiti 設定（beta）"
	MonitoringGraphitiFormDescription = "Graphiti 知識圖譜整合設定"
	MonitoringGraphitiFormName        = "Graphiti（beta）"
	MonitoringGraphitiFormOverview    = `⚠️  BETA FEATURE: This functionality is currently under active development. Please monitor updates for improvements and stability fixes.

Graphiti provides temporal knowledge graph capabilities:
• Entity and relationship extraction
• Semantic memory for AI agents
• Temporal context tracking
• Flow-scoped contextual search

Graphiti reuses credentials configured in the LLM Providers and Embeddings screens. LiteLLM credentials remain in .env. Select the provider preset here; configure its models in ./graphiti/<provider>.yaml.
Supports openai, gemini, litellm, custom.

Choose between embedded instance or external connection.`

	// Deployment types
	MonitoringGraphitiEmbedded = "內嵌堆疊"
	MonitoringGraphitiExternal = "外部服務"
	MonitoringGraphitiDisabled = "已停用"

	// Form fields
	MonitoringGraphitiDeploymentType           = "部署類型"
	MonitoringGraphitiDeploymentTypeDesc       = "選擇 Graphiti 的部署方式"
	MonitoringGraphitiURL                      = "Graphiti 伺服器 URL"
	MonitoringGraphitiURLDesc                  = "Graphiti API 伺服器的位址"
	MonitoringGraphitiTimeout                  = "請求逾時"
	MonitoringGraphitiTimeoutDesc              = "Graphiti 操作的逾時秒數"
	MonitoringGraphitiLLMClientType            = "LLM 供應商預設集"
	MonitoringGraphitiLLMClientTypeDesc        = "供應商預設集；憑證沿用 LLM 供應商設定（openai、gemini、litellm、custom），模型清單來自 ./graphiti/<provider>.yaml"
	MonitoringGraphitiSeparateEmbedding        = "使用獨立嵌入端點"
	MonitoringGraphitiSeparateEmbeddingDesc    = "使用共用的嵌入設定，而非所選 LLM 的憑證"
	MonitoringGraphitiSemaphoreLimit           = "協程上限"
	MonitoringGraphitiSemaphoreLimitDesc       = "Graphiti 輔助協程的最大並行數"
	MonitoringGraphitiLogLevel                 = "日誌等級"
	MonitoringGraphitiLogLevelDesc             = "Graphiti 日誌詳細程度"
	MonitoringGraphitiSearchScope              = "搜尋範圍"
	MonitoringGraphitiSearchScopeDesc          = "使用 flowid 做隔離，或填 all 進行可信任的全域搜尋測試"
	MonitoringGraphitiIngestPolicyRules        = "匯入政策規則"
	MonitoringGraphitiIngestPolicyRulesDesc    = "以 JSON 定義訊息樣式對應到 REJECT、SKIP_LLM 或 PROCESS"
	MonitoringGraphitiIngestPolicyField        = "匯入政策比對欄位"
	MonitoringGraphitiIngestPolicyFieldDesc    = "用於政策比對的訊息欄位"
	MonitoringGraphitiIngestDefaultAction      = "預設匯入動作"
	MonitoringGraphitiIngestDefaultActionDesc  = "當沒有任何匯入政策規則符合時採取的動作"
	MonitoringGraphitiIngestWorkerCount        = "匯入 worker 數量"
	MonitoringGraphitiIngestWorkerCountDesc    = "跨流程的最大並行大量匯入任務數"
	MonitoringGraphitiIngestQueueMaxSize       = "匯入佇列上限"
	MonitoringGraphitiIngestQueueMaxSizeDesc   = "佇列訊息上限；設為 0 表示不限"
	MonitoringGraphitiTaxonomyLayerProfile     = "分類層設定檔"
	MonitoringGraphitiTaxonomyLayerProfileDesc = "啟用的關聯類別，或使用 full／minimal 別名"
	MonitoringGraphitiCPUs                     = "Graphiti CPU 上限"
	MonitoringGraphitiCPUsDesc                 = "Graphiti 容器的 CPU 上限"
	MonitoringGraphitiMemory                   = "Graphiti 記憶體上限"
	MonitoringGraphitiMemoryDesc               = "Graphiti 容器的記憶體上限"
	MonitoringGraphitiNeo4jUser                = "Neo4j 使用者名稱"
	MonitoringGraphitiNeo4jUserDesc            = "存取 Neo4j 資料庫的使用者名稱"
	MonitoringGraphitiNeo4jPassword            = "Neo4j 密碼"
	MonitoringGraphitiNeo4jPasswordDesc        = "存取 Neo4j 資料庫的密碼"
	MonitoringGraphitiNeo4jDatabase            = "Neo4j 資料庫"
	MonitoringGraphitiNeo4jDatabaseDesc        = "Neo4j 資料庫名稱"
	MonitoringGraphitiNeo4jCPUs                = "Neo4j CPU 上限"
	MonitoringGraphitiNeo4jCPUsDesc            = "Neo4j 容器的 CPU 上限"
	MonitoringGraphitiNeo4jMemory              = "Neo4j 記憶體上限"
	MonitoringGraphitiNeo4jMemoryDesc          = "Neo4j 容器的記憶體上限"
	MonitoringGraphitiNeo4jShmSize             = "Neo4j 共享記憶體上限"
	MonitoringGraphitiNeo4jShmSizeDesc         = "/dev/shm 上限；實際用量仍計入 Neo4j 記憶體"

	// Help text
	MonitoringGraphitiModeGuide    = "選擇部署方式：內嵌（本機 Neo4j）、外部（既有 Graphiti）、停用（不使用知識圖譜）"
	MonitoringGraphitiEmbeddedHelp = `⚠️  BETA: This feature is under active development. Monitor updates for improvements.

Embedded deploys complete Graphiti stack:
• Neo4j graph database
• Graphiti API service
• Automatic entity extraction from agent interactions
• Temporal relationship tracking
• Private knowledge graph on your server

Prerequisites:
• Configure provider credentials in LLM Providers (or .env for LiteLLM)
• Configure shared embeddings before enabling separate embeddings
• Edit models and call parameters in ./graphiti/<provider>.yaml

Resource requirements:
• Defaults reserve up to 2GB for Graphiti and 4GB for Neo4j
• Neo4j UI: http://localhost:7474
• Graphiti API: http://localhost:8000
• Automatic setup and maintenance

Keep flowid search scope for isolation. Use a bounded ingest queue on memory-constrained hosts. Semaphore/workers control throughput; Graphiti and Neo4j CPU/memory limits must grow with concurrent flows.

Fine tuning for retries, combined extraction, anchors, telemetry, and diagnostics remains available in .env.

Best for: Teams wanting knowledge graph capabilities with full data control and privacy.`
	MonitoringGraphitiExternalHelp = `⚠️  BETA: This feature is under active development. Monitor updates for improvements.

External connects to your existing Graphiti server:

• No local infrastructure needed
• Managed updates and maintenance
• Shared knowledge graph across teams
• Data stored on external provider

Setup requirements:
• Graphiti server URL and access
• Network connectivity required
• Provider, model, embedding, extraction, and graph settings are configured on the external server

Best for: Teams using existing Graphiti deployments or cloud services.`
	MonitoringGraphitiDisabledHelp = `Graphiti is disabled. You will not have:

• Temporal knowledge graph
• Entity and relationship extraction
• Semantic memory for AI agents
• Flow-scoped graph recall
• Advanced contextual search

PentAGI continues using its primary vector memory when Graphiti is disabled.`
)

// Observability Integration constants
const (
	MonitoringObservabilityFormTitle       = "可觀測性設定"
	MonitoringObservabilityFormDescription = "監控與可觀測性堆疊設定"
	MonitoringObservabilityFormName        = "可觀測性"
	MonitoringObservabilityFormOverview    = `Observability stack includes:
• Grafana dashboards for visualization
• VictoriaMetrics for time-series data
• Jaeger for distributed tracing
• Loki for log aggregation
• OpenTelemetry for data collection

Monitor PentAGI performance and system health.`

	// Deployment types
	MonitoringObservabilityEmbedded = "內嵌堆疊"
	MonitoringObservabilityExternal = "外部收集器"
	MonitoringObservabilityDisabled = "已停用"

	// Form fields
	MonitoringObservabilityDeploymentType     = "部署類型"
	MonitoringObservabilityDeploymentTypeDesc = "選擇監控的部署方式"
	MonitoringObservabilityOTelHost           = "OpenTelemetry 主機"
	MonitoringObservabilityOTelHostDesc       = "外部 OpenTelemetry 收集器的位址"

	// embedded listen fields
	MonitoringObservabilityGrafanaListenIP        = "Grafana 監聽 IP"
	MonitoringObservabilityGrafanaListenIPDesc    = "Docker 連接埠對應使用的綁定位址（例如 0.0.0.0 表示對所有網路介面開放）"
	MonitoringObservabilityGrafanaListenPort      = "Grafana 監聽連接埠"
	MonitoringObservabilityGrafanaListenPortDesc  = "Grafana Web 介面對外暴露的 TCP 連接埠"
	MonitoringObservabilityOTelGrpcListenIP       = "OTel gRPC 監聽 IP"
	MonitoringObservabilityOTelGrpcListenIPDesc   = "Docker 連接埠對應使用的綁定位址（例如 0.0.0.0 表示對所有網路介面開放）"
	MonitoringObservabilityOTelGrpcListenPort     = "OTel gRPC 監聽連接埠"
	MonitoringObservabilityOTelGrpcListenPortDesc = "OTel gRPC 接收器對外暴露的 TCP 連接埠"
	MonitoringObservabilityOTelHttpListenIP       = "OTel HTTP 監聽 IP"
	MonitoringObservabilityOTelHttpListenIPDesc   = "Docker 連接埠對應使用的綁定位址（例如 0.0.0.0 表示對所有網路介面開放）"
	MonitoringObservabilityOTelHttpListenPort     = "OTel HTTP 監聽連接埠"
	MonitoringObservabilityOTelHttpListenPortDesc = "OTel HTTP 接收器對外暴露的 TCP 連接埠"

	// Help text
	MonitoringObservabilityModeGuide    = "選擇監控方式：內嵌（完整堆疊）、外部（既有基礎設施）、停用（不監控）"
	MonitoringObservabilityEmbeddedHelp = `Embedded deploys complete monitoring:
• Grafana dashboards and alerting
• VictoriaMetrics time-series database
• Jaeger distributed tracing UI
• Loki log aggregation system
• ClickHouse analytical database
• Node Exporter + cAdvisor metrics
• OpenTelemetry data collection

Auto-instrumented components with
pre-built dashboards for system health,
performance analysis, and debugging.

Resource requirements:
• ~1.5GB RAM, 3GB disk space minimum
• Grafana UI: http://localhost:3000
• Profiling: http://localhost:7777

Best for: Complete system visibility,
troubleshooting, and performance tuning.`
	MonitoringObservabilityExternalHelp = `External sends telemetry to your existing monitoring infrastructure:

• OTLP protocol over HTTP/2 (no TLS)
• Your collector must support:
  - OTLP HTTP receiver (port 4318)
  - OTLP gRPC receiver (port 8148)
  - tls: insecure: true setting
• Sends metrics, traces, and logs
• Compatible with enterprise platforms:
  Datadog, New Relic, Splunk, etc.

OTEL_HOST example:
your-collector:4318

Collector config requirement:
tls: insecure: true

Best for: Organizations with existing
monitoring infrastructure or centralized
observability platforms.`
	MonitoringObservabilityDisabledHelp = `Observability is disabled. You will not have:

• System performance monitoring
• Distributed request tracing
• Structured log aggregation
• Resource usage analytics
• Error tracking and alerting
• Performance bottleneck analysis

Consider enabling for production use
to monitor system health, debug issues,
and optimize performance effectively.`
)

// Summarizer Screen
const (
	SummarizerTitle       = "摘要器設定"
	SummarizerDescription = "啟用對話摘要，以降低 LLM 成本並改善上下文管理"
	SummarizerName        = "摘要器"
	SummarizerOverview    = `Optimize context usage, reduce LLM costs, and match your model capabilities.

When to adjust summarization:
• High token costs: Reduce context size (4K-12K vs 22K+ tokens)
• "上下文過長" errors: Configure for your model's limits
• Poor conversation flow: Increase context retention for quality
• Different model types: Short-context vs long-context model tuning

General Summarization: Maximum cost control and precision tuning for research/analysis tasks
Assistant Summarization: Optimal conversation quality with intelligent context management for interactive sessions

Quick wins:
• Cost reduction: Use General, reduce Recent Sections to 1-2
• Context errors: Match limits to your model (8K/32K/128K)
• Quality priority: Use Assistant with increased limits`

	SummarizerTypeGeneralName = "一般內容摘要"
	SummarizerTypeGeneralDesc = "用於對話上下文管理的全域摘要設定"

	SummarizerTypeGeneralInfo = `Choose this for maximum cost control and short-context model compatibility.

Perfect when you need:
• Aggressive cost reduction: Fine-tune every parameter for minimal token usage
• Short-context models (8K-32K): Precise limits to avoid overflow errors
• Research/analysis tasks: Controlled compression without losing key data
• Custom QA handling: Full control over question-answer pair processing

Typical results:
• 40-70% cost reduction vs default settings
• 4K-12K token contexts (vs 22K+ in Assistant mode)
• Better performance on GPT-3.5, Claude Instant, smaller models
• Precise control over conversation memory vs fresh context balance

Best practices:
• Start with 1-2 Recent Sections for maximum savings
• Enable Size Management for automatic overflow protection
• Disable QA compression only for critical reasoning tasks`

	SummarizerTypeAssistantName = "助理對話摘要"
	SummarizerTypeAssistantDesc = "AI 助理上下文的專用摘要設定"

	SummarizerTypeAssistantInfo = `Choose this for optimal conversation quality and dialogue continuity.

Perfect when you need:
• Extended reasoning chains: Maintain context for complex multi-step thinking
• High-quality conversations: Preserve dialogue flow and assistant personality
• Long-context models (64K+): Leverage full model capabilities efficiently
• Interactive sessions: Better memory of user preferences and conversation history

Typical results:
• 8K-40K token contexts with intelligent compression
• Superior conversation continuity vs manual settings
• Automatic context optimization for reasoning tasks
• Balanced cost vs quality (3x more context than General mode)

Best practices:
• Use default settings for most scenarios - they're pre-optimized
• Increase Recent Sections only for very complex tasks
• Monitor context usage - costs scale with token count
• Perfect for GPT-4, Claude, and other large context models`
)

// Summarizer Form Screen
const (
	SummarizerFormGeneralTitle   = "一般摘要設定"
	SummarizerFormAssistantTitle = "助理摘要設定"
	SummarizerFormDescription    = "設定 %s 選項"

	// Field Labels and Descriptions
	SummarizerFormPreserveLast     = "大小管理"
	SummarizerFormPreserveLastDesc = "控制末段壓縮。啟用：區段會壓到 LastSecBytes 以內（上下文較小）；停用：區段自由成長（上下文較大）"

	SummarizerFormUseQA     = "問答摘要"
	SummarizerFormUseQADesc = "當問答內容總量超過 MaxQABytes 或 MaxQASections 限制時，啟用問答配對壓縮"

	SummarizerFormSumHumanInQA     = "壓縮使用者訊息"
	SummarizerFormSumHumanInQADesc = "問答壓縮時包含使用者訊息。停用：保留使用者原文（多數情況建議停用）"

	SummarizerFormLastSecBytes     = "區段大小上限"
	SummarizerFormLastSecBytesDesc = "啟用大小管理時，每個近期區段的位元組上限。數值越大：細節越多、Token 用量越高"

	SummarizerFormMaxBPBytes     = "回應大小上限"
	SummarizerFormMaxBPBytesDesc = "單則 AI 回應在壓縮前的位元組上限，避免單一過長回應佔滿上下文"

	SummarizerFormMaxQASections     = "問答區段上限"
	SummarizerFormMaxQASectionsDesc = "觸發問答壓縮前的問答區段上限，與 MaxQABytes 共同控制問答記憶的總量"

	SummarizerFormMaxQABytes     = "問答記憶總量"
	SummarizerFormMaxQABytesDesc = "所有問答區段合計的位元組上限。超過時（連同 MaxQASections）會觸發問答壓縮以符合限制"

	SummarizerFormKeepQASections     = "近期區段"
	SummarizerFormKeepQASectionsDesc = "保留不壓縮的最近對話區段數。這是影響上下文大小的主要參數"

	// Enhanced Help Text - General (common principles)
	SummarizerFormGeneralHelp = `Context estimation: 4K-22K tokens (typical), up to 94K (maximum settings).

Key relationships:
• Recent Sections: Most critical - each +1 adds ~1.5-9K tokens
• Size Management OFF: 2-3x larger context (less compression)
• Section/Response Limits: Control individual component sizes
• QA Memory: Manages total conversation history when limits exceeded

Parameter interactions:
• QA compression activates when BOTH MaxQABytes AND MaxQASections exceeded
• Size Management disabled → sections can grow 2x larger than limits
• Response Limit prevents single large outputs from dominating context
• User message compression (SummHumanInQA) saves 5% but loses original phrasing

Reduce for smaller models:
• Recent Sections: 1-2 (vs 3+ default)
• Section Limit: 25-35KB (vs 50KB+)
• Disable Size Management for simple conversations

Common mistakes:
• Setting Recent Sections too high (main cause of context overflow)
• Enabling Size Management with very low Section Limits (over-compression)
• Mismatched QA limits (high bytes + low sections = ineffective)

Current algorithm compresses older content while preserving recent context quality.`

	// Enhanced Help Text - Assistant specific (interactive conversations)
	SummarizerFormAssistantHelp = `Optimized for interactive conversations requiring context continuity.

Default tuning (3 Recent Sections, 75KB limits):
• Typical range: 8K-40K tokens
• Good for: Extended dialogues, reasoning chains, context-dependent tasks
• Models: Works well with 32K+ context models

Adjustments by model type:
• Short context (≤16K): Recent Sections=1-2, Section Limit=45KB
• Long context (128K+): Can increase Recent Sections=5-7
• High-frequency chat: Reduce Recent Sections=2 for faster responses

Advanced tuning:
• QA Memory 200KB+ for document analysis conversations
• Response Limit 24-32KB for detailed technical responses
• Keep User Messages uncompressed (SummHumanInQA=false) for better context

Performance optimization:
• Each Recent Section ≈ 9-18KB in assistant mode
• Size Management reduces growth by ~20% but may lose detail
• QA compression triggers less often due to larger default limits

Size Management enabled by default - maintains conversation flow while preventing context overflow.
Monitor actual token usage and adjust Recent Sections first, then limits.`

	// Context size estimation
	SummarizerContextEstimatedSize    = "預估上下文大小：%s\\n%s"
	SummarizerContextTokenRange       = "約 %s 個 Token"
	SummarizerContextTokenRangeMinMax = "約 %s-%s 個 Token"
	SummarizerContextRequires256K     = "需要 256K 以上上下文的模型"
	SummarizerContextRequires128K     = "需要 128K 以上上下文的模型"
	SummarizerContextRequires64K      = "需要 64K 以上上下文的模型"
	SummarizerContextRequires32K      = "需要 32K 以上上下文的模型"
	SummarizerContextRequires16K      = "需要 16K 以上上下文的模型"
	SummarizerContextFitsIn8K         = "適用於 8K 以上上下文的模型"
)

// Tools screen strings
const (
	ToolsTitle       = "工具設定"
	ToolsDescription = "以額外工具與選項強化代理能力"
	ToolsName        = "工具"
	ToolsOverview    = `Configure additional tools and capabilities for AI agents.
Each tool can be enabled and configured according to your requirements.

Available settings:
• Human-in-the-loop - Enable user interaction during testing
• AI Agents Settings - Configure global behavior for AI agents
• Search Engines - Configure external search providers
• Scraper - Web content extraction and analysis
• Graphiti (beta) - Temporal knowledge graph for semantic memory
• Docker - Container environment configuration`
)

// Server Settings screen strings
const (
	ServerSettingsFormTitle       = "伺服器設定"
	ServerSettingsFormDescription = "設定 PentAGI 伺服器的網路存取與對外路由"
	ServerSettingsFormName        = "伺服器設定"
	ServerSettingsFormOverview    = `• Network binding - control which interface and port PentAGI listens on
• Public URL - external address and optional base path used in redirects
• CORS - allowed origins for browser access
• Proxy - HTTP/HTTPS proxy for outbound traffic to LLM/search providers
• SSL directory - custom certificates directory containing server.crt and server.key (PEM)
• Data directory - persistent storage for agent artifacts and flow workspaces`

	// Field labels and descriptions
	ServerSettingsLicenseKey     = "授權金鑰"
	ServerSettingsLicenseKeyDesc = "PentAGI 授權金鑰，格式為 XXXX-XXXX-XXXX-XXXX"

	ServerSettingsUpdateStrategy = "更新策略"
	// The three values are spelled out because the field is free text: the
	// validator refuses anything else on save, and a user who has to guess the
	// vocabulary will find that out the hard way.
	ServerSettingsUpdateStrategyDesc = "更新來源通道：preview（預設，跟隨 " +
		"the channels you already use), stable (published releases only) or nightly (newest build, " +
		"released or not)"

	ServerSettingsUpdateServerHost     = "更新伺服器主機"
	ServerSettingsUpdateServerHostDesc = "PentAGI 雲端 API 端點（host[:port]）；留空則使用 update.pentagi.com"

	ServerSettingsSupportServerHost     = "支援伺服器主機"
	ServerSettingsSupportServerHostDesc = "PentAGI 支援 API 端點（host[:port]）；留空則使用 support.pentagi.com"

	ToolsDockerInsideHost     = "Worker Docker Daemon 主機"
	ToolsDockerInsideHostDesc = "提供給 worker 容器的 Docker daemon 端點（例如 tcp://dind:2376）；留空則維持 socket 掛載"

	ToolsDockerInsideTLSVerify     = "Worker Docker TLS 驗證"
	ToolsDockerInsideTLSVerifyDesc = "為 worker 容器的 Docker 連線啟用 TLS 驗證"

	ToolsDockerInsideCertPath     = "Worker Docker 憑證路徑"
	ToolsDockerInsideCertPathDesc = "WORKER 節點上的 TLS 憑證目錄，以唯讀方式掛載到 worker 容器"

	ToolsDockerInsideHostHelp = `Docker daemon endpoint handed to worker containers when Docker Access is enabled.

It is injected into each sandbox as DOCKER_HOST, so the Docker CLI inside the container talks to this daemon.

Setting it also STOPS the host socket from being auto-detected and bind-mounted into sandboxes: an agent then reaches only the daemon you designated, not the one running PentAGI itself. An explicitly configured Docker Socket still takes precedence and is mounted as before.

Leave empty to keep the historical behaviour (auto-detected socket mount).

Examples:
• tcp://dind:2376
• tcp://10.0.0.5:2375
• unix:///var/run/docker.sock`

	ToolsDockerInsideTLSVerifyHelp = `Enables TLS verification for the worker container's connection to the Docker daemon.

Injected into each sandbox as DOCKER_TLS_VERIFY. Use it together with Worker Docker Daemon Host and Worker Docker Certificate Path when the daemon requires mutual TLS.

Leave disabled for a plain (non-TLS) endpoint.`

	ToolsDockerInsideCertPathHelp = `Directory holding the TLS client material (ca.pem, cert.pem, key.pem) used by worker containers.

Injected into each sandbox as DOCKER_CERT_PATH and bind-mounted read-only at the SAME path, so the value resolves identically inside the container.

IMPORTANT: this path is resolved on the WORKER NODE — the machine whose Docker daemon creates sandbox containers — not on the machine running this installer. It is therefore not validated here; make sure it exists on the worker node.

Examples:
• /etc/docker/certs
• /opt/pentagi/docker/ssl`

	ServerSettingsTenantID     = "租戶 ID"
	ServerSettingsTenantIDDesc = "同一主機部署多個實例時可用的命名空間（例如 acme、team_alpha）"

	ServerSettingsPprofAddr     = "pprof 監聽位址"
	ServerSettingsPprofAddrDesc = "選填的 Go pprof 綁定位址；留空表示停用（例如 :7777、127.0.0.1:7778）"

	ServerSettingsHost     = "伺服器主機（監聽 IP）"
	ServerSettingsHostDesc = "Docker 連接埠對應使用的綁定位址（例如 0.0.0.0 表示對所有網路介面開放）"

	ServerSettingsPort     = "伺服器連接埠（監聽連接埠）"
	ServerSettingsPortDesc = "PentAGI Web 介面對外暴露的 TCP 連接埠"

	ServerSettingsPublicURL     = "公開 URL"
	ServerSettingsPublicURLDesc = "用於重新導向與連結的公開基底 URL（支援子路徑，例如 https://example.com/pentagi/）"

	ServerSettingsCORSOrigins     = "CORS 來源"
	ServerSettingsCORSOriginsDesc = "以逗號分隔的允許來源清單（例如 https://localhost:8443,https://localhost）"

	ServerSettingsTrustedProxies     = "可信任代理"
	ServerSettingsTrustedProxiesDesc = "以逗號分隔、其 X-Forwarded-For 可被信任的代理 IP／CIDR（留空表示全不信任）"

	ServerSettingsProxyURL     = "HTTP/HTTPS 代理"
	ServerSettingsProxyURLDesc = "後端 LLM、嵌入與搜尋請求使用的代理（不含沙箱、擷取器、Graphiti）"

	ServerSettingsProxyUsername     = "代理使用者名稱"
	ServerSettingsProxyUsernameDesc = "代理驗證使用者名稱（選填）"
	ServerSettingsProxyPassword     = "代理密碼"
	ServerSettingsProxyPasswordDesc = "代理驗證密碼（選填）"

	ServerSettingsHTTPClientTimeout       = "HTTP 用戶端逾時"
	ServerSettingsHTTPClientTimeoutDesc   = "外部 API 呼叫的逾時秒數（LLM 供應商、搜尋引擎等）"
	ServerSettingsTerminalToolTimeout     = "終端機工具逾時"
	ServerSettingsTerminalToolTimeoutDesc = "終端機指令的預設逾時秒數（0 或負數表示最多 3 小時）"

	ServerSettingsExternalSSLCAPath     = "自訂 CA 憑證路徑"
	ServerSettingsExternalSSLCAPathDesc = "容器內自訂根 CA 憑證的路徑（例如 /opt/pentagi/ssl/ca-bundle.pem）"

	ServerSettingsExternalSSLInsecure     = "略過 SSL 驗證"
	ServerSettingsExternalSSLInsecureDesc = "停用 SSL/TLS 憑證驗證（僅建議在測試自簽憑證時使用）"

	ServerSettingsSSLDir     = "SSL 目錄"
	ServerSettingsSSLDirDesc = "包含 PEM 格式 server.crt 與 server.key 的目錄（server.crt 可含完整憑證鏈）"

	ServerSettingsDataDir     = "資料目錄"
	ServerSettingsDataDirDesc = "代理產生檔案存放的目錄；內含 flow-N 子目錄，會掛載為 worker 容器中的 /work"

	ServerSettingsCookieSigningSalt     = "Cookie 簽章 Salt"
	ServerSettingsCookieSigningSaltDesc = "用於簽署 Cookie 的密鑰（請保密）"

	ServerSettingsDatabaseExtensionsSchema     = "資料庫擴充結構"
	ServerSettingsDatabaseExtensionsSchemaDesc = "設定租戶 ID 時，存放共用擴充的結構（例如 public、Supabase 的 extensions）"

	ServerSettingsDatabaseSearchPathViaOptions     = "以選項搜尋路徑"
	ServerSettingsDatabaseSearchPathViaOptionsDesc = "在啟動參數 options 中傳送租戶 search_path（部分連線池需要）"

	// Hints for fields overview
	ServerSettingsLicenseKeyHint                   = "授權金鑰"
	ServerSettingsTenantIDHint                     = "租戶 ID"
	ServerSettingsPprofAddrHint                    = "pprof 位址"
	ServerSettingsHostHint                         = "監聽 IP"
	ServerSettingsPortHint                         = "監聽連接埠"
	ServerSettingsPublicURLHint                    = "公開 URL"
	ServerSettingsCORSOriginsHint                  = "CORS 來源"
	ServerSettingsProxyURLHint                     = "代理 URL"
	ServerSettingsProxyUsernameHint                = "代理使用者名稱"
	ServerSettingsProxyPasswordHint                = "代理密碼"
	ServerSettingsHTTPClientTimeoutHint            = "HTTP 逾時"
	ServerSettingsTerminalToolTimeoutHint          = "終端機逾時"
	ServerSettingsExternalSSLCAPathHint            = "自訂 CA 路徑"
	ServerSettingsExternalSSLInsecureHint          = "略過 SSL 驗證"
	ServerSettingsSSLDirHint                       = "SSL 目錄"
	ServerSettingsDataDirHint                      = "資料目錄"
	ServerSettingsDatabaseExtensionsSchemaHint     = "擴充結構"
	ServerSettingsDatabaseSearchPathViaOptionsHint = "以選項搜尋路徑"

	// Help texts per-field
	ServerSettingsGeneralHelp = `PentAGI exposes its web UI via Docker with configurable host and port.

Public URL must reflect how users reach the server. If using a subpath (e.g., /pentagi/), include it here. CORS controls browser access from specified origins. Proxy affects outbound traffic to LLM/search providers and other external services used by Tools.

SSL directory allows providing custom certificates. When set, server will use server.crt and server.key from that directory. Data directory stores artifacts and working files for flows.`

	ServerSettingsLicenseKeyHelp = `PentAGI License Key in format of XXXX-XXXX-XXXX-XXXX. It's used to communicate with PentAGI Cloud API.`

	ServerSettingsTenantIDHelp = `Optional identifier that namespaces PostgreSQL schema, data directory, Docker objects, Graphiti group ids, auth cookies and telemetry when several PentAGI instances share one host and the same backing services.

Leave empty for a single-instance deployment (default). When set, the value must match ^[a-z][a-z0-9_]{0,31}$ — lowercase letter first, then lowercase letters, digits or underscores, max 32 characters. Hyphens are not allowed.

Examples:
• acme
• team_alpha
• staging01`

	ServerSettingsPprofAddrHelp = `Optional listen address for the Go pprof HTTP endpoint used for CPU/memory profiling.

Leave empty to keep pprof disabled (default, recommended for production). When set, the value must be a host:port pair such as :7777 or 127.0.0.1:7778.

Instances that share a host network namespace need distinct addresses — a port is not a string namespace and cannot be derived from TENANT_ID.

Examples:
• :7777
• 127.0.0.1:7778
• 0.0.0.0:7780`

	ServerSettingsHostHelp = `Bind address for published port in docker-compose mapping.

Examples:
• 127.0.0.1 — local-only access
• 0.0.0.0 — expose on all interfaces`

	ServerSettingsPortHelp = `External port for PentAGI UI. Must be available on the host. Example: 8443.`

	ServerSettingsPublicURLHelp = `Set the public base URL used in redirects and links.

Examples:
• http://localhost:8443
• https://example.com/
• https://example.com/pentagi/ (with base path)`

	ServerSettingsCORSOriginsHelp = `Comma-separated allowed origins for browser access.`

	ServerSettingsTrustedProxiesHelp = `Comma-separated list of proxy IP addresses or CIDR ranges whose X-Forwarded-For / X-Real-IP headers PentAGI will trust when determining a client's real IP.

Set this only when PentAGI sits behind a reverse proxy or load balancer you control. Leave empty to trust none. Example: 10.0.0.0/8,172.16.0.0/12`

	ServerSettingsProxyURLHelp = `HTTP or HTTPS proxy for every backend request to LLM, embedding and search providers, local endpoints included, and for update checks; the installer's own update and download calls use it too. Sandbox containers, the scraper, Graphiti and Docker API communication do not use it.`

	ServerSettingsHTTPClientTimeoutHelp = `Timeout in seconds for all external HTTP/HTTPS API calls including:
• LLM provider requests (OpenAI, Anthropic, Bedrock, etc.)
• Search engine queries (Google, Tavily, Perplexity, etc.)
• External tool integrations
• Embedding generation requests

Default: 600 seconds (10 minutes)
Setting to 0 disables timeout (not recommended in production)
Too low values may cause legitimate long-running requests to fail.`

	ServerSettingsTerminalToolTimeoutHelp = `Default timeout in seconds applied when an agent requests timeout=0 or a negative timeout value.

This affects commands executed through the isolated terminal container, including scanners and CLI-based utilities.

Default: 1200 seconds (20 minutes)
Allowed range: 1–10800 seconds (up to 3 hours)
Values <= 0 or above 10800 are clamped to the maximum (10800 s = 3 hours); agents are never allowed to run indefinitely.
Explicit timeout values provided by the tool call override this default when they are within the 1–10800 s range.`

	ServerSettingsExternalSSLCAPathHelp = `Path to custom CA certificate file (PEM format) inside the container.

Must point to /opt/pentagi/ssl/ directory, which is mounted from pentagi-ssl volume on the host.

Examples:
• /opt/pentagi/ssl/ca-bundle.pem
• /opt/pentagi/ssl/corporate-ca.pem

File can contain multiple root and intermediate certificates.`

	ServerSettingsExternalSSLInsecureHelp = `Disable SSL/TLS certificate validation for connections to LLM providers and external services.

⚠ WARNING: Use only for testing with self-signed certificates. Never enable in production.

When enabled, all certificate validation is bypassed, making connections vulnerable to man-in-the-middle attacks.`

	ServerSettingsSSLDirHelp = `Path to directory with server.crt and server.key in PEM format. server.crt may include fullchain. Overrides default generated certificate behavior.`

	ServerSettingsDataDirHelp = `Host directory for persistent data. PentAGI stores agent artifacts under flow-N subdirectories, which map to /work inside worker containers.`

	ServerSettingsCookieSigningSaltHelp = `Secret salt used to sign cookies. Keep it private.`

	ServerSettingsDatabaseExtensionsSchemaHelp = `PostgreSQL schema that holds the shared extensions (vector, pg_trgm) every tenant must reach through its search_path.

Only used when Tenant ID is set; leave empty for the default "public", which is where a stock PostgreSQL install keeps them. Set it when your database follows another convention — Supabase installs its extensions into "extensions", and startup aborts with a message naming the schema it found if this does not match.

Examples:
• public
• extensions`

	ServerSettingsDatabaseSearchPathViaOptionsHelp = `Sends the tenant search_path as options=--search_path=<value> instead of a bare search_path connection parameter.

Only used when Tenant ID is set. Keep it false for a direct PostgreSQL connection. Enable it when connecting through a pooler that forwards the "options" startup parameter but drops an unrecognized bare search_path — reported to be the case for some versions of Supabase's Supavisor. It is not guaranteed to work: startup fails with a clear schema-mismatch error if the value never reaches the backend.

Values: true, false`
)

// Human-in-the-loop screen strings
const (
	// AI Agents Settings screen strings
	ToolsAIAgentsSettingsFormTitle       = "AI 代理設定"
	ToolsAIAgentsSettingsFormDescription = "設定 AI 代理的全域行為"
	ToolsAIAgentsSettingsFormName        = "AI 代理設定"
	ToolsAIAgentsSettingsFormOverview    = `This section configures global behavior of AI agents across PentAGI.

Basic Settings:
• Enable User Interaction: allow agents to request user input when needed
• Use Multi-Agent Mode: enable assistant to orchestrate multiple specialized agents

Execution Monitoring (⚠️  BETA):
• Enable Execution Monitoring: automatic mentor supervision for pattern analysis
• Same Tool Call Threshold: consecutive identical tool calls before mentor review
• Total Tool Call Threshold: total tool calls before mentor review

Tool Call Limits:
• Max Tool Calls (General Agents): prevent runaway executions for Assistant, Primary Agent, Pentester, Coder, Installer
• Max Tool Calls (Limited Agents): prevent runaway executions for Searcher, Enricher, Memorist, etc.

Task Planning (⚠️  BETA):
• Enable Task Planning: generate structured execution plans for specialist agents

⚠️  BETA features are under active development. Enable for testing only.`

	// field labels and descriptions
	ToolsAIAgentsSettingHumanInTheLoop          = "啟用使用者互動"
	ToolsAIAgentsSettingHumanInTheLoopDesc      = "允許代理在需要時向使用者提問"
	ToolsAIAgentsSettingUseAgents               = "使用多代理模式"
	ToolsAIAgentsSettingUseAgentsDesc           = "讓助理能調度多個專業代理"
	ToolsAIAgentsSettingExecutionMonitor        = "啟用執行監控（beta）"
	ToolsAIAgentsSettingExecutionMonitorDesc    = "自動請導師進行執行模式分析"
	ToolsAIAgentsSettingSameToolLimit           = "相同工具呼叫門檻"
	ToolsAIAgentsSettingSameToolLimitDesc       = "連續相同工具呼叫達到幾次後交由導師審查"
	ToolsAIAgentsSettingTotalToolLimit          = "工具呼叫總數門檻"
	ToolsAIAgentsSettingTotalToolLimitDesc      = "交由導師審查前的工具呼叫總數"
	ToolsAIAgentsSettingMaxGeneralToolCalls     = "最大工具呼叫數（一般代理）"
	ToolsAIAgentsSettingMaxGeneralToolCallsDesc = "助理、主代理、滲透測試、程式、安裝代理的最大工具呼叫數"
	ToolsAIAgentsSettingMaxLimitedToolCalls     = "最大工具呼叫數（受限代理）"
	ToolsAIAgentsSettingMaxLimitedToolCallsDesc = "搜尋、補充、記憶等代理的最大工具呼叫數"
	ToolsAIAgentsSettingTaskPlanning            = "啟用任務規劃（beta）"
	ToolsAIAgentsSettingTaskPlanningDesc        = "為專業代理產生結構化的執行計畫"
	ToolsAIAgentsSettingLLMFallbackProvider     = "LLM 備援供應商"
	ToolsAIAgentsSettingLLMFallbackProviderDesc = "主要供應商失敗時改用哪個供應商重試（留空表示不重試）"

	// help content
	ToolsAIAgentsSettingsHelp = `AI Agents Settings define how agents collaborate, interact with users, and handle execution control.

Basic Settings:
• Enable User Interaction: allow agents to ask for user input when needed
• Use Multi-Agent Mode: enable assistant to orchestrate specialized agents for complex tasks

Execution Monitoring (⚠️  BETA):
Automatically invokes adviser (mentor) to analyze execution patterns, detect loops, suggest alternative strategies, and prevent agents from fixating on single approach. Thresholds: consecutive identical calls (default: 5) and total calls (default: 10).

Task Planning (⚠️  BETA):
Generates 3-7 step execution plans before specialist agents begin work. Prevents scope creep and improves success rates. Works best when adviser uses enhanced configuration (stronger model or maximum reasoning mode).

Tool Call Limits (always active):
Hard limits prevent infinite loops: General agents default 100, Limited agents default 20. Works independently from beta features.

OPEN SOURCE MODELS < 32B (Qwen3.5-27B, DeepSeek-V3, Llama-3.1-70B):
✓ ENABLE both beta features - ESSENTIAL for quality results
✓ Testing shows 2x improvement in result quality vs. baseline
✓ Configure adviser with enhanced settings for best performance
✓ Ideal for air-gapped deployments with local LLM inference

Performance: 2-3x increase in tokens/time, 2x improvement in quality for models < 32B.

⚠️  BETA WARNING: Features under active development. Recommended for open source models < 32B despite beta status. For cloud APIs with larger models, keep disabled.

Note: Changes require service restart.`
)

// Search Engines screen strings
const (
	ToolsSearchEnginesFormTitle       = "搜尋引擎設定"
	ToolsSearchEnginesFormDescription = "設定搜尋引擎，讓 AI 代理在測試期間蒐集情報"
	ToolsSearchEnginesFormName        = "搜尋引擎"
	ToolsSearchEnginesFormOverview    = `Available search engines:
• DuckDuckGo - Free search engine (no API key required)
• Sploitus - Security exploits and vulnerabilities database (no API key required)
• Perplexity - AI-powered search with reasoning
• Tavily - Search API for AI applications
• Firecrawl - Web search with full-page scraping (self-hosting supported)
• Traversaal - Web scraping and search
• Google Search - Requires API key and Custom Search Engine ID
• Searxng - Internet metasearch engine
• Internal Analytics Engine - Optional browser-based fallback (scrape + summarize, off by default, no API key required)

Get API keys from:
• Perplexity: https://www.perplexity.ai/
• Tavily: https://tavily.com/
• Firecrawl: https://www.firecrawl.dev/
• Traversaal: https://traversaal.ai/
• Google: https://developers.google.com/custom-search/v1/introduction`

	ToolsSearchEnginesDuckDuckGo               = "DuckDuckGo 搜尋"
	ToolsSearchEnginesDuckDuckGoDesc           = "啟用 DuckDuckGo 搜尋（不需要 API 金鑰）"
	ToolsSearchEnginesDuckDuckGoRegion         = "DuckDuckGo 區域"
	ToolsSearchEnginesDuckDuckGoRegionDesc     = "DuckDuckGo 區域代碼（例如 us-en、uk-en、cn-zh）"
	ToolsSearchEnginesDuckDuckGoSafeSearch     = "DuckDuckGo 安全搜尋"
	ToolsSearchEnginesDuckDuckGoSafeSearchDesc = "DuckDuckGo 安全搜尋（strict、moderate、off）"
	ToolsSearchEnginesDuckDuckGoTimeRange      = "DuckDuckGo 時間範圍"
	ToolsSearchEnginesDuckDuckGoTimeRangeDesc  = "DuckDuckGo 時間範圍（d：日、w：週、m：月、y：年）"
	ToolsSearchEnginesSploitus                 = "Sploitus 搜尋"
	ToolsSearchEnginesSploitusDesc             = "啟用 Sploitus 漏洞與 exploit 搜尋（不需要 API 金鑰）"
	ToolsSearchEnginesPerplexityKey            = "Perplexity API 金鑰"
	ToolsSearchEnginesPerplexityKeyDesc        = "Perplexity AI 搜尋的 API 金鑰"
	ToolsSearchEnginesPerplexityTimeout        = "Perplexity 逾時"
	ToolsSearchEnginesPerplexityTimeoutDesc    = "Perplexity 請求逾時秒數（留空為 120）"
	ToolsSearchEnginesTavilyKey                = "Tavily API 金鑰"
	ToolsSearchEnginesTavilyKeyDesc            = "Tavily 搜尋服務的 API 金鑰"
	ToolsSearchEnginesFirecrawlKey             = "Firecrawl API 金鑰"
	ToolsSearchEnginesFirecrawlKeyDesc         = "Firecrawl 搜尋服務的 API 金鑰"
	ToolsSearchEnginesFirecrawlURL             = "Firecrawl API URL"
	ToolsSearchEnginesFirecrawlURLDesc         = "Firecrawl API 基底 URL（雲端服務請留空；自架時填入）"
	ToolsSearchEnginesTraversaalKey            = "Traversaal API 金鑰"
	ToolsSearchEnginesTraversaalKeyDesc        = "Traversaal 網頁擷取的 API 金鑰"
	ToolsSearchEnginesGoogleKey                = "Google 搜尋 API 金鑰"
	ToolsSearchEnginesGoogleKeyDesc            = "Google 自訂搜尋 API 金鑰"
	ToolsSearchEnginesGoogleCX                 = "Google 搜尋引擎 ID"
	ToolsSearchEnginesGoogleCXDesc             = "Google 自訂搜尋引擎 ID"
	ToolsSearchEnginesGoogleLR                 = "Google 語言限制"
	ToolsSearchEnginesGoogleLRDesc             = "Google 搜尋引擎語言限制（例如 lang_en、lang_cn 等）"
	ToolsSearchEnginesSearxngURL               = "Searxng 搜尋 URL"
	ToolsSearchEnginesSearxngURLDesc           = "Searxng 搜尋引擎 URL"
	ToolsSearchEnginesSearxngCategories        = "Searxng 搜尋分類"
	ToolsSearchEnginesSearxngCategoriesDesc    = "Searxng 搜尋分類（例如 general、it、web、news、technology、science、health、other）"
	ToolsSearchEnginesSearxngLanguage          = "Searxng 搜尋語言"
	ToolsSearchEnginesSearxngLanguageDesc      = "Searxng 搜尋語言（en、ch、fr、de、it、es、pt、ru、zh；留空表示全部）"
	ToolsSearchEnginesSearxngSafeSearch        = "Searxng 安全搜尋"
	ToolsSearchEnginesSearxngSafeSearchDesc    = "Searxng 安全搜尋（0：關閉、1：中等、2：嚴格）"
	ToolsSearchEnginesSearxngTimeRange         = "Searxng 時間範圍"
	ToolsSearchEnginesSearxngTimeRangeDesc     = "Searxng 時間範圍（day、month、year）"
	ToolsSearchEnginesSearxngTimeout           = "Searxng 逾時"
	ToolsSearchEnginesSearxngTimeoutDesc       = "Searxng 請求逾時秒數"
	ToolsSearchEnginesInternalEnabled          = "內部分析引擎"
	ToolsSearchEnginesInternalEnabledDesc      = "為問答／研究查詢啟用內建瀏覽器分析備援（不需 API 金鑰；會擷取並摘要頁面，因此需要已設定的擷取器與至少一個可用的連結搜尋引擎，例如 DuckDuckGo 或 Google）"
	ToolsSearchEnginesInternalMaxSites         = "內部引擎站台數上限"
	ToolsSearchEnginesInternalMaxSitesDesc     = "每次查詢可抓取並摘要的頁數上限"
	ToolsSearchEnginesInternalMaxSiteBytes     = "內部引擎單站台位元組上限"
	ToolsSearchEnginesInternalMaxSiteBytesDesc = "每頁讀取的 Markdown 位元組上限（超過即截斷）"
)

// Scraper screen strings
const (
	ToolsScraperFormTitle       = "擷取器設定"
	ToolsScraperFormDescription = "設定網頁擷取服務"
	ToolsScraperFormName        = "擷取器"
	ToolsScraperFormOverview    = `Web scraper service for content extraction and analysis using vxcontrol/scraper Docker image.

Modes:
• Embedded - Run local scraper container (recommended)
• External - Use external scraper services
• Disabled - No web scraping capabilities

Docker image: https://hub.docker.com/r/vxcontrol/scraper

The scraper supports:
• Public URL access for external links
• Private URL access for internal/local links
• Content extraction and analysis
• Multiple output formats`

	ToolsScraperModeTitle                 = "擷取器模式"
	ToolsScraperModeDesc                  = "選擇擷取服務的運作方式"
	ToolsScraperEmbedded                  = "內嵌容器"
	ToolsScraperExternal                  = "外部服務"
	ToolsScraperDisabled                  = "已停用"
	ToolsScraperPublicURL                 = "公開擷取服務 URL"
	ToolsScraperPublicURLDesc             = "擷取公開／外部網站所用的 URL。留空則使用與私有 URL 相同的值。"
	ToolsScraperPublicURLEmbeddedDesc     = "內嵌擷取器的 URL（選填覆寫）。留空則使用與私有 URL 相同的值。"
	ToolsScraperPrivateURL                = "私有擷取服務 URL"
	ToolsScraperPrivateURLDesc            = "擷取私有／內部網站所用的 URL"
	ToolsScraperPublicUsername            = "公開 URL 使用者名稱"
	ToolsScraperPublicUsernameDesc        = "公開擷取服務的使用者名稱"
	ToolsScraperPublicPassword            = "公開 URL 密碼"
	ToolsScraperPublicPasswordDesc        = "公開擷取服務的密碼"
	ToolsScraperPrivateUsername           = "私有 URL 使用者名稱"
	ToolsScraperPrivateUsernameDesc       = "私有擷取服務的使用者名稱"
	ToolsScraperPrivatePassword           = "私有 URL 密碼"
	ToolsScraperPrivatePasswordDesc       = "私有擷取服務的密碼"
	ToolsScraperLocalUsername             = "本機 URL 使用者名稱"
	ToolsScraperLocalUsernameDesc         = "內嵌擷取服務的使用者名稱"
	ToolsScraperLocalPassword             = "本機 URL 密碼"
	ToolsScraperLocalPasswordDesc         = "內嵌擷取服務的密碼"
	ToolsScraperMaxConcurrentSessions     = "最大並行工作階段數"
	ToolsScraperMaxConcurrentSessionsDesc = "並行擷取工作階段的上限"
	ToolsScraperEmbeddedHelp              = "內嵌模式會執行本機擷取容器，可同時存取公開與私有資源。預設設定使用 https://someuser:somepass@scraper/。"
	ToolsScraperExternalHelp              = "外部模式使用獨立的擷取服務。可依需求為公開與私有存取設定不同 URL。"
	ToolsScraperDisabledHelp              = "擷取器已停用，將無法使用網頁內容擷取與分析功能。"
)

// Docker Environment screen strings
const (
	ToolsDockerFormTitle       = "Docker 環境設定"
	ToolsDockerFormDescription = "設定 worker 容器使用的 Docker 環境"
	ToolsDockerFormName        = "Docker 環境"
	ToolsDockerFormOverview    = `• Worker Isolation - Containers provide security boundaries for tasks
• Network Capabilities - Enable privileged network operations for pentesting
• Container Management - Control how workers access Docker daemon
• Storage Configuration - Define workspace and artifact storage
• Image Selection - Set default images for different task types

Critical for penetration testing workflows requiring network scanning, custom tools, and secure task isolation.`

	// General help text
	ToolsDockerGeneralHelp = `Each AI agent task runs in an isolated Docker container with two ports allocated per flow from the 2000-port window that starts at DOCKER_PORTS_BASE (28000-29999 by default). Worker containers are created on-demand from default images or agent-selected ones.

Basic setup requires enabling capabilities: Docker Access allows spawning additional containers for specialized tools, while Network Admin grants low-level network permissions essential for scanning tools like nmap.

Storage operates via Docker volumes by default, or host directories when Work Directory is specified. Connection settings control the Docker daemon location - local socket for standard setups, or remote TCP with TLS for distributed environments.

Default images serve as fallbacks: general tasks use standard images, while security testing defaults to pentesting-focused containers. Public IP enables reverse shell attacks by providing workers with a reachable address for target callbacks. Usually it's a local interface address of the host machine with Docker daemon running for the workers containers.

Configuration combines based on scenario: enable both capabilities for full pentesting, use Work Directory for persistent artifacts, or configure remote connection for isolated Docker environments.`

	// Container capabilities
	ToolsDockerInside       = "Docker 存取"
	ToolsDockerInsideDesc   = "允許 worker 管理 Docker 容器"
	ToolsDockerNetAdmin     = "網路管理"
	ToolsDockerNetAdminDesc = "為 nmap 等網路掃描工具授予 NET_ADMIN 能力"

	// Connection settings
	ToolsDockerSocket       = "Docker Socket"
	ToolsDockerSocketDesc   = "主機檔案系統上的 Docker Socket 路徑"
	ToolsDockerNetwork      = "Docker 網路"
	ToolsDockerNetworkDesc  = "worker 容器使用的自訂網路名稱，或填 'host' 直接使用主機網路"
	ToolsDockerPublicIP     = "公開 IP 位址"
	ToolsDockerPublicIPDesc = "OOB 攻擊中反向連線使用的公開 IP"

	// Storage configuration
	ToolsDockerWorkDir     = "工作目錄"
	ToolsDockerWorkDirDesc = "worker 檔案系統的主機目錄（預設使用 Docker 資料卷）"

	// Default images
	ToolsDockerDefaultImage               = "預設映像"
	ToolsDockerDefaultImageDesc           = "一般任務使用的預設 Docker 映像"
	ToolsDockerDefaultImageForPentest     = "滲透測試映像"
	ToolsDockerDefaultImageForPentestDesc = "安全測試任務使用的預設 Docker 映像"
	ToolsDockerDefaultImageForTest        = "沙箱自我測試映像"
	ToolsDockerDefaultImageForTestDesc    = "啟動時執行沙箱隔離檢查所用的小型映像"
	ToolsDockerImageSelectionMode         = "映像選擇模式"
	ToolsDockerImageSelectionModeDesc     = "worker 映像的選擇方式：'llm'（由代理挑選）或 'fixed'（固定使用滲透測試映像）"
	ToolsDockerAllowedImages              = "允許的映像"
	ToolsDockerAllowedImagesDesc          = "以逗號分隔的允許清單，供代理從中挑選（留空表示任何映像皆可）"
	ToolsDockerInsidePolicyTests          = "沙箱隔離自我測試"
	ToolsDockerInsidePolicyTestsDesc      = "啟動時驗證沙箱 daemon 為獨立隔離，並拒絕逃逸到主機"

	// TLS connection settings (optional)
	ToolsDockerHost          = "Docker 主機"
	ToolsDockerHostDesc      = "Docker daemon 連線（unix:// 或 tcp://）"
	ToolsDockerTLSVerify     = "TLS 驗證"
	ToolsDockerTLSVerifyDesc = "為 Docker 連線啟用 TLS 驗證"
	ToolsDockerCertPath      = "TLS 憑證"
	ToolsDockerCertPathDesc  = "包含 ca.pem、cert.pem、key.pem 的目錄"

	// Help content for specific configurations
	ToolsDockerInsideHelp = `Docker Access enables workers to spawn additional containers for specialized tools and environments. Required when tasks need custom software not available in default images.

When enabled, workers can pull and run any Docker image, providing maximum flexibility for complex testing scenarios.`

	ToolsDockerNetAdminHelp = `Network Admin capability allows workers to perform low-level network operations essential for penetration testing.

Required for:
• Network scanning with nmap, masscan
• Custom packet crafting
• Network interface manipulation
• Raw socket operations

Critical for comprehensive security assessments.`

	ToolsDockerSocketHelp = `Docker Socket path defines how workers access the Docker daemon. Use only file path to the socket file. Used with Docker Access to enable container management.

For enhanced security, consider using docker-in-docker (DinD) instead of exposing the main Docker daemon directly to workers.
When using DinD, use the path to the Docker socket file of the DinD container which binded to the host filesystem.

Example: /var/run/docker.sock`

	ToolsDockerNetworkHelp = `Docker Network controls network isolation mode for worker containers:

Bridge Mode (custom network name):
• Isolated communication between containers
• Port forwarding from container to host
• Enhanced security boundaries
• Network-based monitoring and filtering
• Recommended for most use cases

Host Mode (value: 'host'):
• Direct access to host network interfaces
• No port forwarding - ports bind directly to host
• Required for raw packet manipulation
• Advanced network testing capabilities
• Lower isolation - use with caution

Examples:
• 'pentagi-network' - creates isolated bridge network
• 'host' - enables direct host network access

Security Note: Host network mode reduces container isolation. Only use when necessary for advanced penetration testing tasks requiring direct network stack access.`

	ToolsDockerPublicIPHelp = `Public IP Address enables out-of-band (OOB) attack techniques by providing workers with a reachable address for reverse connections.

Each flow's worker gets two ports from the 2000-port window that starts at DOCKER_PORTS_BASE (28000-29999 by default), published on this IP for receiving callbacks from exploited targets. The default 0.0.0.0 publishes them on every interface of the host.

By default agents will try to get public address from the services api.ipify.org, ipinfo.io/ip or ifconfig.me.`

	ToolsDockerWorkDirHelp = `Work Directory specifies host filesystem location for worker storage. When set, replaces default Docker volumes with host directory mounts.

Benefits:
• Persistent storage across restarts
• Direct file system access
• Easier artifact management
• Custom backup strategies

By default uses Docker dedicated volume per worker container.

Example: /path/to/workdir/`

	ToolsDockerDefaultImageHelp = `Default Image provides fallback for workers when task requirements don't specify a particular container image.

Should contain basic utilities and tools for general-purpose tasks. Default: debian:latest`

	ToolsDockerDefaultImageForPentestHelp = `Pentesting Image serves as default for security testing tasks. Should include comprehensive security tools and utilities.

Recommended images include Kali Linux, Parrot Security, or custom security-focused containers. Default: vxcontrol/kali-linux`

	ToolsDockerDefaultImageForTestHelp = `Sandbox Self-Test Image is the small image the startup isolation check runs in and pulls into the sandbox to prove an agent could. Kept small because the check pays for it on every start.

Only used when Docker Access and its Isolation Self-Test are enabled. Default: vxcontrol/kali-linux:test`

	ToolsDockerImageSelectionModeHelp = `Image Selection Mode controls how a worker's container image is chosen.

• llm — the agent picks the image for each task (default)
• fixed — always use the Pentesting Image, skipping the model's choice

Use 'fixed' to pin every worker to one vetted image.`

	ToolsDockerAllowedImagesHelp = `Allowed Images is a comma-separated allow-list the agent may pick from when Image Selection Mode is 'llm'. A choice outside the list falls back to the Pentesting Image.

Leave empty to allow any image. Example: debian:latest,vxcontrol/kali-linux`

	ToolsDockerInsidePolicyTestsHelp = `Sandbox Isolation Self-Test checks at startup that the sandbox Docker daemon is separate from PentAGI's own and refuses host-escape requests, before any agent is given Docker access.

Only meaningful when Docker Access is enabled. Recommended on any deployment where agents reach a Docker daemon.`

	ToolsDockerHostHelp = `Docker Host uses for start primary worker containers and overrides default Docker daemon connection. Supports Unix sockets and TCP connections.

Examples:
• unix:///var/run/docker.sock (local)
• tcp://docker-host:2376 (remote)

Enable TLS for remote connections.`

	ToolsDockerTLSVerifyHelp = `TLS Verification secures Docker daemon connections over TCP. Strongly recommended for remote Docker hosts.

Requires valid certificates in the specified certificate directory.`

	ToolsDockerCertPathHelp = `TLS Certificates directory must contain:
• ca.pem - Certificate Authority
• cert.pem - Client certificate
• key.pem - Private key

Required for secure remote Docker connections when using TLS to manage worker containers.

Example: /path/to/certs`
)

// Embedder form strings
const (
	EmbedderFormTitle       = "嵌入模型設定"
	EmbedderFormDescription = "設定文字向量化，用於語意搜尋與知識儲存"
	EmbedderFormName        = "嵌入模型"
	EmbedderFormOverview    = `Text embeddings convert documents into vectors for semantic search and knowledge storage.
Different providers offer various models with different capabilities and pricing.

Choose carefully as changing providers requires reindexing all stored data.`

	EmbedderFormProvider     = "嵌入供應商"
	EmbedderFormProviderDesc = "選擇文字向量化供應商。嵌入向量用於語意搜尋與知識儲存。"

	EmbedderFormURL     = "API 端點 URL"
	EmbedderFormURLDesc = "自訂 API 端點（留空則使用預設值）"

	EmbedderFormAPIKey     = "API 金鑰"
	EmbedderFormAPIKeyDesc = "此供應商的驗證金鑰（Ollama 不需要）"

	EmbedderFormModel     = "模型名稱"
	EmbedderFormModelDesc = "指定使用的嵌入模型（留空則用供應商預設）"

	EmbedderFormBatchSize     = "批次大小"
	EmbedderFormBatchSizeDesc = "單一批次處理的文件數量（1-1000）"

	EmbedderFormStripNewLines     = "移除換行"
	EmbedderFormStripNewLinesDesc = "嵌入前移除文字換行（true/false）"

	EmbedderFormMaxTextBytes     = "文字位元組上限"
	EmbedderFormMaxTextBytesDesc = "送往嵌入 API 的每個文字區塊位元組上限（例如 8192）"

	// Provider names and descriptions
	EmbedderProviderDefault         = "預設（OpenAI）"
	EmbedderProviderDefaultDesc     = "使用 OpenAI 嵌入，金鑰取自 LLM 供應商設定"
	EmbedderProviderOpenAI          = "OpenAI"
	EmbedderProviderOpenAIDesc      = "OpenAI 文字嵌入 API（text-embedding-3-small、ada-002）"
	EmbedderProviderOllama          = "Ollama"
	EmbedderProviderOllamaDesc      = "本機 Ollama 伺服器，用於開源嵌入模型"
	EmbedderProviderMistral         = "Mistral"
	EmbedderProviderMistralDesc     = "Mistral AI 嵌入模型"
	EmbedderProviderJina            = "Jina"
	EmbedderProviderJinaDesc        = "Jina AI 嵌入 API"
	EmbedderProviderHuggingFace     = "HuggingFace"
	EmbedderProviderHuggingFaceDesc = "HuggingFace 推論 API（用於嵌入模型）"
	EmbedderProviderGoogleAI        = "Google AI"
	EmbedderProviderGoogleAIDesc    = "Google AI 嵌入模型（gemini-embedding-001）"
	EmbedderProviderVoyageAI        = "VoyageAI"
	EmbedderProviderVoyageAIDesc    = "VoyageAI 嵌入 API"
	EmbedderProviderDisabled        = "已停用"
	EmbedderProviderDisabledDesc    = "完全停用嵌入功能"

	// Provider-specific placeholders and help
	EmbedderURLPlaceholderOpenAI      = "https://api.openai.com/v1"
	EmbedderURLPlaceholderOllama      = "http://localhost:11434"
	EmbedderURLPlaceholderMistral     = "https://api.mistral.ai/v1"
	EmbedderURLPlaceholderJina        = "https://api.jina.ai/v1"
	EmbedderURLPlaceholderHuggingFace = "https://api-inference.huggingface.co"
	EmbedderURLPlaceholderGoogleAI    = "不支援 —— 使用預設端點"
	EmbedderURLPlaceholderVoyageAI    = "https://api.voyageai.com/v1"

	EmbedderAPIKeyPlaceholderOllama      = "本機模型不需要"
	EmbedderAPIKeyPlaceholderMistral     = "Mistral API 金鑰"
	EmbedderAPIKeyPlaceholderJina        = "Jina API 金鑰"
	EmbedderAPIKeyPlaceholderHuggingFace = "HuggingFace API 金鑰"
	EmbedderAPIKeyPlaceholderGoogleAI    = "Google AI API 金鑰"
	EmbedderAPIKeyPlaceholderVoyageAI    = "VoyageAI API 金鑰"
	EmbedderAPIKeyPlaceholderDefault     = "此供應商的 API 金鑰"

	EmbedderModelPlaceholderOpenAI      = "text-embedding-3-small"
	EmbedderModelPlaceholderOllama      = "nomic-embed-text"
	EmbedderModelPlaceholderMistral     = "mistral-embed"
	EmbedderModelPlaceholderJina        = "jina-embeddings-v2-base-en"
	EmbedderModelPlaceholderHuggingFace = "sentence-transformers/all-MiniLM-L6-v2"
	EmbedderModelPlaceholderGoogleAI    = "gemini-embedding-001"
	EmbedderModelPlaceholderVoyageAI    = "voyage-2"
	EmbedderModelPlaceholderDefault     = "模型名稱"

	// Provider IDs for internal use
	EmbedderProviderIDDefault     = "default"
	EmbedderProviderIDOpenAI      = "openai"
	EmbedderProviderIDOllama      = "ollama"
	EmbedderProviderIDMistral     = "mistral"
	EmbedderProviderIDJina        = "jina"
	EmbedderProviderIDHuggingFace = "huggingface"
	EmbedderProviderIDGoogleAI    = "googleai"
	EmbedderProviderIDVoyageAI    = "voyageai"
	EmbedderProviderIDDisabled    = "none"

	EmbedderHelpGeneral = `Embeddings convert text into vectors for semantic search and knowledge storage. This enables PentAGI to understand meaning rather than just keywords, making search results more relevant and intelligent.

Key benefits:
• Find documents by meaning, not exact words
• Build a smart knowledge base from pentesting results
• Enable AI agents to locate relevant information quickly
• Support advanced reasoning with contextual data

Choose Ollama for completely local processing - your data never leaves your infrastructure. Other providers offer cloud-based processing with different model capabilities and pricing.

Configure carefully as changing providers requires rebuilding the entire knowledge base.`

	EmbedderHelpAttentionPrefix = "重要："
	EmbedderHelpAttention       = `Different embedding providers create incompatible vectors. Changing providers or models will break existing semantic search.

You must flush or reindex your entire knowledge base using the etester utility:
• Run 'etester flush' to clear old embeddings
• Run 'etester reindex' to rebuild with new provider
• This process can take significant time for large datasets`

	EmbedderHelpAttentionSuffix = `Only change providers if absolutely necessary.`

	// Provider help texts
	EmbedderHelpOpenAIKeyRule = `When the embedder has no API Key, it uses the OpenAI key from LLM Providers, and only with the OpenAI server set there: its API Endpoint URL must be empty or name that server. An embedder API Key goes to the API Endpoint URL, or to api.openai.com when the URL is empty. Any other endpoint needs its own API Key; a local server that checks no key accepts any non-empty value.`

	EmbedderHelpDefault = `Default mode uses OpenAI embeddings with the server and API key configured for OpenAI in LLM Providers, and needs nothing more if OpenAI is set up there.

An API Endpoint URL or API Key left from another embedding provider still applies in this mode; choose OpenAI to see or clear them.

` + EmbedderHelpOpenAIKeyRule

	EmbedderHelpOpenAI = `Direct OpenAI API access for embedding generation.

Get your API key from:
https://platform.openai.com/api-keys

` + EmbedderHelpOpenAIKeyRule + `

Recommended models:
• text-embedding-3-small (cost-effective, 1536 dimensions)
• text-embedding-3-large (highest quality, 3072 dimensions)
• text-embedding-ada-002 (legacy, still supported)`

	EmbedderHelpOllama = `Local Ollama server for open-source embedding models.

Popular embedding models:
• nomic-embed-text (recommended, 768 dimensions)
• mxbai-embed-large (large model, 1024 dimensions)
• snowflake-arctic-embed (multilingual support)

Install Ollama from:
https://ollama.com/

Start with: ollama pull nomic-embed-text`

	EmbedderHelpMistral = `Mistral AI embedding models via API.

Get your API key from:
https://console.mistral.ai/

Leave the Model Name empty to use mistral-embed.`

	EmbedderHelpJina = `Jina AI embedding API with specialized models.

Get your API key from:
https://jina.ai/

Recommended models:
• jina-embeddings-v2-base-en (general purpose, 768 dimensions)
• jina-embeddings-v2-small-en (lightweight, 512 dimensions)
• jina-embeddings-v2-base-code (code-specific embeddings)`

	EmbedderHelpHuggingFace = `HuggingFace Inference API for open-source models.

Get your API key from:
https://huggingface.co/settings/tokens

Popular models:
• sentence-transformers/all-MiniLM-L6-v2 (384 dimensions)
• sentence-transformers/all-mpnet-base-v2 (768 dimensions)
• intfloat/e5-large-v2 (1024 dimensions)`

	EmbedderHelpGoogleAI = `Google AI embedding models (Gemini).

Get your API key from:
https://aistudio.google.com/app/apikey

Available models:
• gemini-embedding-001 (latest model, 768 dimensions)
• text-embedding-004 (legacy Vertex AI model)

Uses Google's fixed endpoint - URL configuration not supported.`

	EmbedderHelpVoyageAI = `VoyageAI embedding API optimized for retrieval.

Get your API key from:
https://www.voyageai.com/

Recommended models:
• voyage-2 (general purpose, 1024 dimensions)
• voyage-large-2 (highest quality, 1536 dimensions)
• voyage-code-2 (code embeddings, 1536 dimensions)`

	EmbedderHelpDisabled = `Disables all embedding functionality.

This will:
• Disable semantic search capabilities
• Turn off knowledge storage vectorization
• Reduce memory and computational requirements

Only recommended if embeddings are not needed for your use case.`
)

// Development and Mock Screen constants
const (
	MockScreenTitle       = "開發畫面"
	MockScreenDescription = "此畫面仍在開發中"
)

// Apply Changes screen constants
const (
	ApplyChangesFormTitle       = "套用設定變更"
	ApplyChangesFormName        = "套用變更"
	ApplyChangesFormDescription = "檢視並套用你的設定變更"

	// Apply Changes overview and help
	ApplyChangesFormOverview = `This screen allows you to review all pending configuration changes and apply them to your PentAGI installation.

When you apply changes, the system will:
• Save all modified environment variables to the .env file
• Restart affected services with the new configuration
• Install additional components if needed`

	// Apply Changes status messages
	ApplyChangesNotStarted     = "設定變更已就緒，可以套用"
	ApplyChangesInProgress     = "正在套用設定變更…\\n"
	ApplyChangesCompleted      = "設定變更已成功套用\\n"
	ApplyChangesFailed         = "套用設定變更失敗"
	ApplyChangesResetCompleted = "設定變更已成功重設\\n"

	ApplyChangesTerminalIsNotInitialized = "終端機尚未初始化"

	// Apply Changes instructions
	ApplyChangesInstructions = `Press Enter to begin applying the configuration changes.`

	ApplyChangesNoChanges = "沒有待套用的設定變更"

	// Apply Changes installation status
	ApplyChangesInstallNotFound = `PentAGI is not currently installed on this system.

The following actions will be performed:
• Docker environment setup and validation
• Creation of docker-compose.yml file
• Installation and startup of PentAGI core services`

	ApplyChangesInstallFoundLangfuse      = `• Installation of Langfuse observability stack (docker-compose-langfuse.yml)`
	ApplyChangesInstallFoundObservability = `• Installation of comprehensive observability stack with Grafana, VictoriaMetrics, and Jaeger (docker-compose-observability.yml)`

	ApplyChangesUpdateFound = `PentAGI is currently installed on this system.

The following actions will be performed:
• Update environment variables in .env file
• Recreate and restart affected Docker containers
• Apply new configuration to running services`

	// Apply Changes warnings and notes
	ApplyChangesWarningCritical = "⚠️ 偵測到重大變更 —— 服務將會重新啟動"
	ApplyChangesWarningSecrets  = "🔒 偵測到機密值 —— 將以安全方式儲存"
	ApplyChangesNoteBackup      = "💾 變更前會先備份目前設定"
	ApplyChangesNoteTime        = "⏱️ 此流程所需的時間依所選元件而定，通常不到一分鐘"

	// Apply Changes progress messages
	ApplyChangesStageValidation = "正在驗證環境與相依套件…"
	ApplyChangesStageBackup     = "正在建立設定備份…"
	ApplyChangesStageEnvFile    = "正在更新環境設定檔…"
	ApplyChangesStageCompose    = "正在產生 Docker Compose 檔案…"
	ApplyChangesStageDocker     = "正在管理 Docker 容器…"
	ApplyChangesStageServices   = "正在啟動服務…"
	ApplyChangesStageComplete   = "設定變更已成功套用"

	// Apply Changes change list headers
	ApplyChangesChangesTitle  = "待套用的設定變更"
	ApplyChangesChangesCount  = "變更總數：%d"
	ApplyChangesChangesMasked = "（基於安全考量已隱藏）"
	ApplyChangesChangesEmpty  = "沒有需要套用的變更"

	// Apply Changes help content
	ApplyChangesHelpTitle   = "正在套用設定變更"
	ApplyChangesHelpContent = `Be sure to check the current configuration before applying changes.`
)

// apply changes integrity prompt
const (
	ApplyChangesIntegrityPromptTitle   = "檔案完整性檢查"
	ApplyChangesIntegrityPromptMessage = "偵測到過期的檔案。\\n要將它們更新到最新版本嗎？"
	ApplyChangesIntegrityOutdatedList  = "過期的檔案：\\n%s\\n確認更新？（y/n）"
	ApplyChangesIntegrityChecking      = "正在收集檔案完整性資訊…"
	ApplyChangesIntegrityNoOutdated    = "沒有找到過期的檔案，繼續套用。"
)

// Maintenance Screen constants
const (
	MaintenanceTitle       = "系統維護"
	MaintenanceDescription = "管理 PentAGI 服務並執行維護操作"
	MaintenanceName        = "維護"
	MaintenanceOverview    = `Perform system maintenance operations for PentAGI.

Available operations depend on the current system state and will only be shown when applicable.

Operations include:
• Service lifecycle management (Start/Stop/Restart)
• Component updates and downloads
• System reset and cleanup
• Container and image management

Each operation will provide real-time status updates and confirmation when required.`

	// Maintenance menu items
	MaintenanceStartPentagi            = "啟動 PentAGI"
	MaintenanceStartPentagiDesc        = "啟動所有已設定的 PentAGI 服務"
	MaintenanceStopPentagi             = "停止 PentAGI"
	MaintenanceStopPentagiDesc         = "停止所有執行中的 PentAGI 服務"
	MaintenanceRestartPentagi          = "重新啟動 PentAGI"
	MaintenanceRestartPentagiDesc      = "重新啟動所有 PentAGI 服務"
	MaintenanceDownloadWorkerImage     = "下載 worker 映像"
	MaintenanceDownloadWorkerImageDesc = "下載 worker 任務使用的滲透測試容器映像"
	MaintenanceUpdateWorkerImage       = "更新 worker 映像"
	MaintenanceUpdateWorkerImageDesc   = "將滲透測試容器映像更新到最新版本"
	MaintenanceUpdatePentagi           = "更新 PentAGI"
	MaintenanceUpdatePentagiDesc       = "將 PentAGI 更新到最新版本"
	MaintenanceUpdateInstaller         = "更新安裝程式"
	MaintenanceUpdateInstallerDesc     = "下載並驗證最新的安裝程式版本"
	MaintenanceFactoryReset            = "出廠重設"
	MaintenanceFactoryResetDesc        = "將 PentAGI 重設為出廠預設值"
	MaintenanceRemovePentagi           = "移除 PentAGI"
	MaintenanceRemovePentagiDesc       = "移除 PentAGI 容器但保留資料"
	MaintenancePurgePentagi            = "清除 PentAGI"
	MaintenancePurgePentagiDesc        = "完全移除 PentAGI，包含所有資料"
	MaintenanceResetPassword           = "重設管理員密碼"
	MaintenanceResetPasswordDesc       = "重設 PentAGI 的管理員密碼"
)

// Reset Password Screen constants
const (
	ResetPasswordFormTitle       = "重設管理員密碼"
	ResetPasswordFormDescription = "重設 PentAGI 的管理員密碼"
	ResetPasswordFormName        = "重設密碼"
	ResetPasswordFormOverview    = `Reset the password for the default administrator account (admin@pentagi.com).

This operation requires PentAGI to be running and will update the password in the PostgreSQL database.

Enter your new password twice to confirm and press Enter to apply the change.

Password requirements:
• 16+ characters, or 8-15 with a lowercase letter, an uppercase letter, a digit and one of !@#$&*
• Both password fields must match`

	// Form fields
	ResetPasswordNewPassword         = "新密碼"
	ResetPasswordNewPasswordDesc     = "輸入新的管理員密碼"
	ResetPasswordConfirmPassword     = "確認密碼"
	ResetPasswordConfirmPasswordDesc = "再次輸入新密碼以確認"

	// Status messages
	ResetPasswordNotAvailable = "必須在 PentAGI 執行中才能重設密碼"
	ResetPasswordAvailable    = "可以重設密碼"
	ResetPasswordInProgress   = "正在重設密碼…"
	ResetPasswordSuccess      = "密碼已成功重設"
	ResetPasswordErrorPrefix  = "錯誤："

	// Validation errors
	ResetPasswordErrorEmptyPassword = "密碼不能為空"
	ResetPasswordErrorWeakPassword  = "密碼須為 16 個字元以上，或 8-15 個字元且包含小寫、大寫、數字與 !@#$&* 其中之一"
	ResetPasswordErrorLongPassword  = "密碼不得超過 72 位元組"
	ResetPasswordErrorMismatch      = "兩次輸入的密碼不一致"

	// Help content
	ResetPasswordHelpContent = `Reset the administrator password for accessing PentAGI.

This operation:
• Updates the password for admin@pentagi.com account
• Sets the user status to 'active'
• Requires PentAGI database to be accessible
• Does not affect other user accounts

The password change takes effect immediately after successful completion.

Enter the same password in both fields and press Enter to confirm the change.`
)

// Processor Operation Form constants
const (
	// Dynamic title templates
	ProcessorOperationFormTitle       = "%s"
	ProcessorOperationFormDescription = "執行 %s 操作"
	ProcessorOperationFormName        = "%s"

	// Common status messages
	ProcessorOperationNotStarted = "已可執行 %s 操作"
	ProcessorOperationInProgress = "正在執行 %s 操作…\\n"
	ProcessorOperationCompleted  = "%s 操作已成功完成\\n"
	ProcessorOperationFailed     = "執行 %s 操作失敗"

	// Confirmation messages
	ProcessorOperationConfirmation = "確定要%s嗎？"
	ProcessorOperationPressEnter   = "按 Enter 以%s"
	ProcessorOperationPressYN      = "按 Y 確認，按 N 取消"
	// Short notice without hotkeys (for static help panel)
	ProcessorOperationRequiresConfirmationShort = "此操作需要確認"
	// Additional terminal messages
	ProcessorOperationCancelled = "操作已取消"
	ProcessorOperationUnknown   = "未知的操作：%s"

	// Operation specific messages
	ProcessorOperationStarting    = "正在啟動服務…"
	ProcessorOperationStopping    = "正在停止服務…"
	ProcessorOperationRestarting  = "正在重新啟動服務…"
	ProcessorOperationDownloading = "正在下載映像…"
	ProcessorOperationUpdating    = "正在更新元件…"
	ProcessorOperationResetting   = "正在重設為出廠預設值…"
	ProcessorOperationRemoving    = "正在移除容器…"
	ProcessorOperationPurging     = "正在清除所有資料…"
	ProcessorOperationInstalling  = "正在安裝 PentAGI 服務…"

	// Help text templates
	ProcessorOperationHelpTitle           = "%s 操作"
	ProcessorOperationHelpContent         = "此操作將%s。"
	ProcessorOperationHelpContentDownload = "此操作將下載 %s 元件。"
	ProcessorOperationHelpContentUpdate   = "此操作將更新 %s 元件。"
	// Generic title/description/builders for dynamic operations
	OperationTitleInstallPentagi    = "安裝 PentAGI"
	OperationDescInstallPentagi     = "安裝並設定 PentAGI 服務"
	OperationTitleDownload          = "下載 %s"
	OperationDescDownloadComponents = "下載 %s 元件"
	OperationTitleUpdate            = "更新 %s"
	OperationDescUpdateToLatest     = "將 %s 更新到最新版本"
	OperationTitleExecute           = "執行 %s"
	OperationDescExecuteOn          = "在 %s 上執行 %s"
	OperationProgressExecuting      = "正在執行 %s…"

	// Terminal not initialized
	ProcessorOperationTerminalNotInitialized = "終端機尚未初始化"
)

// Operation-specific help texts
const (
	ProcessorHelpInstallPentagi = `This will:
• Deploy Docker containers for selected services
• Configure networking and volumes
• Start all enabled services
• Set up monitoring if configured

Installation will use your current configuration settings.`

	ProcessorHelpStartPentagi = `This will:
• Core PentAGI API and web interface
• Configured Langfuse analytics (if enabled)
• Observability stack (if enabled)

Services will be started in the correct dependency order.`

	ProcessorHelpStopPentagi = `This will:
• Gracefully shutdown containers
• Preserve all data and configurations
• Network connections will be closed

You can restart services later without losing any data.`

	ProcessorHelpRestartPentagi = `This will:
• Stop running containers
• Apply any configuration changes
• Start services with fresh state

Useful after configuration updates or to resolve issues.`

	ProcessorHelpDownloadWorkerImage = `This large image (6GB+) contains:
• Kali Linux tools and utilities
• Security testing frameworks
• Network analysis software

Required for pentesting operations.`

	ProcessorHelpUpdateWorkerImage = `This will:
• Pull the latest pentesting image
• Update security tools and frameworks
• Preserve existing worker containers

Note: This is a large download (6GB+).`

	ProcessorHelpUpdatePentagi = `This will:
• Download latest container images
• Perform rolling update of services
• Preserve all data and configurations

Services will be briefly unavailable during update.`

	ProcessorHelpFactoryReset = `⚠️  WARNING: This operation will:
• Remove all containers and networks
• Delete all configuration files
• Clear stored data and volumes
• Restore default settings

This action cannot be undone!`

	ProcessorHelpRemovePentagi = `This will:
• Stop and remove all containers
• Remove Docker networks
• Preserve volumes and data
• Keep configuration files

You can reinstall later without losing data.`

	ProcessorHelpPurgePentagi = `⚠️  WARNING: This will permanently delete:
• All containers and images
• All data volumes
• All configuration files
• All stored results

This action cannot be undone!`
)

// environment variable descriptions (centralized)
const (
	EnvDesc_OPEN_AI_KEY                       = "OpenAI API 金鑰"
	EnvDesc_OPEN_AI_SERVER_URL                = "OpenAI 伺服器 URL"
	EnvDesc_ANTHROPIC_API_KEY                 = "Anthropic API 金鑰"
	EnvDesc_ANTHROPIC_SERVER_URL              = "Anthropic 伺服器 URL"
	EnvDesc_ANTHROPIC_ORGANIZATION_ID         = "Anthropic 組織 ID"
	EnvDesc_ANTHROPIC_WORKSPACE_ID            = "Anthropic 工作區 ID"
	EnvDesc_ANTHROPIC_SERVICE_ACCOUNT_ID      = "Anthropic 服務帳號 ID"
	EnvDesc_ANTHROPIC_IDENTITY_TOKEN          = "Anthropic 身分權杖"
	EnvDesc_ANTHROPIC_IDENTITY_TOKEN_FILE     = "Anthropic 身分權杖檔案"
	EnvDesc_ANTHROPIC_FEDERATION_RULE_ID      = "Anthropic 聯邦規則 ID"
	EnvDesc_GEMINI_API_KEY                    = "Google Gemini API 金鑰"
	EnvDesc_GEMINI_SERVER_URL                 = "Gemini 伺服器 URL"
	EnvDesc_BEDROCK_DEFAULT_AUTH              = "AWS Bedrock 使用預設憑證鏈"
	EnvDesc_BEDROCK_BEARER_TOKEN              = "AWS Bedrock Bearer 權杖"
	EnvDesc_BEDROCK_ACCESS_KEY_ID             = "AWS Bedrock 存取金鑰 ID"
	EnvDesc_BEDROCK_SECRET_ACCESS_KEY         = "AWS Bedrock 私密存取金鑰"
	EnvDesc_BEDROCK_SESSION_TOKEN             = "AWS Bedrock 工作階段權杖"
	EnvDesc_BEDROCK_REGION                    = "AWS Bedrock 區域"
	EnvDesc_BEDROCK_SERVER_URL                = "AWS Bedrock 自訂端點 URL"
	EnvDesc_BEDROCK_CONFIG_PATH               = "AWS Bedrock 容器設定檔路徑"
	EnvDesc_OLLAMA_SERVER_URL                 = "Ollama 伺服器 URL"
	EnvDesc_OLLAMA_SERVER_API_KEY             = "Ollama 伺服器 API 金鑰（雲端）"
	EnvDesc_OLLAMA_SERVER_MODEL               = "Ollama 預設模型"
	EnvDesc_OLLAMA_SERVER_CONFIG_PATH         = "Ollama 容器設定檔路徑"
	EnvDesc_OLLAMA_SERVER_PULL_MODELS_TIMEOUT = "Ollama 模型拉取逾時"
	EnvDesc_OLLAMA_SERVER_PULL_MODELS_ENABLED = "Ollama 自動拉取模型"
	EnvDesc_OLLAMA_SERVER_LOAD_MODELS_ENABLED = "Ollama 載入模型清單"
	EnvDesc_DEEPSEEK_API_KEY                  = "DeepSeek API 金鑰"
	EnvDesc_DEEPSEEK_SERVER_URL               = "DeepSeek 伺服器 URL"
	EnvDesc_DEEPSEEK_PROVIDER                 = "DeepSeek 供應商名稱前綴（供 LiteLLM 使用，例如 'deepseek'）"
	EnvDesc_GLM_API_KEY                       = "GLM API 金鑰"
	EnvDesc_GLM_SERVER_URL                    = "GLM 伺服器 URL"
	EnvDesc_GLM_PROVIDER                      = "GLM 供應商名稱前綴（供 LiteLLM 使用，例如 'zai'）"
	EnvDesc_KIMI_API_KEY                      = "Kimi API 金鑰"
	EnvDesc_KIMI_SERVER_URL                   = "Kimi 伺服器 URL"
	EnvDesc_KIMI_PROVIDER                     = "Kimi 供應商名稱前綴（供 LiteLLM 使用，例如 'moonshot'）"
	EnvDesc_QWEN_API_KEY                      = "Qwen API 金鑰"
	EnvDesc_QWEN_SERVER_URL                   = "Qwen 伺服器 URL"
	EnvDesc_QWEN_PROVIDER                     = "Qwen 供應商名稱前綴（供 LiteLLM 使用，例如 'dashscope'）"
	EnvDesc_MINIMAX_API_KEY                   = "MiniMax API 金鑰"
	EnvDesc_MINIMAX_SERVER_URL                = "MiniMax 伺服器 URL"
	EnvDesc_MINIMAX_PROVIDER                  = "MiniMax 供應商名稱前綴（供 LiteLLM 使用，例如 'minimax'）"
	EnvDesc_MISTRAL_API_KEY                   = "Mistral API 金鑰"
	EnvDesc_MISTRAL_SERVER_URL                = "Mistral 伺服器 URL"
	EnvDesc_MISTRAL_PROVIDER                  = "Mistral 供應商名稱前綴（供 LiteLLM 使用，例如 'mistral'）"
	EnvDesc_XAI_API_KEY                       = "xAI API 金鑰"
	EnvDesc_XAI_SERVER_URL                    = "xAI 伺服器 URL"
	EnvDesc_XAI_PROVIDER                      = "xAI 供應商名稱前綴（供 LiteLLM 使用，例如 'xai'）"
	EnvDesc_LLM_SERVER_URL                    = "自訂 LLM 伺服器 URL"
	EnvDesc_LLM_SERVER_KEY                    = "自訂 LLM API 金鑰"
	EnvDesc_LLM_SERVER_MODEL                  = "自訂 LLM 模型"
	EnvDesc_LLM_SERVER_CONFIG_PATH            = "自訂 LLM 容器設定檔路徑"
	EnvDesc_LLM_SERVER_PRESERVE_REASONING     = "自訂 LLM 保留推理內容"
	EnvDesc_LLM_SERVER_API_TYPE               = "自訂 LLM 位址慣例"
	EnvDesc_LLM_SERVER_API_VERSION            = "自訂 LLM Azure API 版本"
	EnvDesc_LLM_SERVER_PROVIDER               = "自訂 LLM 供應商名稱"
	EnvDesc_LLM_FALLBACK_PROVIDER             = "LLM 備援供應商"

	EnvDesc_LANGFUSE_LISTEN_IP   = "Langfuse 監聽 IP"
	EnvDesc_LANGFUSE_LISTEN_PORT = "Langfuse 監聽連接埠"
	EnvDesc_LANGFUSE_BASE_URL    = "Langfuse 基底 URL"
	EnvDesc_LANGFUSE_PROJECT_ID  = "Langfuse 專案 ID"
	EnvDesc_LANGFUSE_PUBLIC_KEY  = "Langfuse 公開金鑰"
	EnvDesc_LANGFUSE_SECRET_KEY  = "Langfuse 私密金鑰"

	// langfuse init variables
	EnvDesc_LANGFUSE_INIT_PROJECT_ID         = "Langfuse 初始化專案 ID"
	EnvDesc_LANGFUSE_INIT_PROJECT_PUBLIC_KEY = "Langfuse 初始化專案公開金鑰"
	EnvDesc_LANGFUSE_INIT_PROJECT_SECRET_KEY = "Langfuse 初始化專案私密金鑰"
	EnvDesc_LANGFUSE_INIT_USER_EMAIL         = "Langfuse 初始化使用者電子郵件"
	EnvDesc_LANGFUSE_INIT_USER_NAME          = "Langfuse 初始化使用者名稱"
	EnvDesc_LANGFUSE_INIT_USER_PASSWORD      = "Langfuse 初始化使用者密碼"

	EnvDesc_LANGFUSE_OTEL_EXPORTER_OTLP_ENDPOINT = "OpenTelemetry 匯出器使用的 Langfuse OTLP 端點"

	EnvDesc_GRAFANA_LISTEN_IP     = "Grafana 監聽 IP"
	EnvDesc_GRAFANA_LISTEN_PORT   = "Grafana 監聽連接埠"
	EnvDesc_OTEL_GRPC_LISTEN_IP   = "OTel gRPC 監聽 IP"
	EnvDesc_OTEL_GRPC_LISTEN_PORT = "OTel gRPC 監聽連接埠"
	EnvDesc_OTEL_HTTP_LISTEN_IP   = "OTel HTTP 監聽 IP"
	EnvDesc_OTEL_HTTP_LISTEN_PORT = "OTel HTTP 監聽連接埠"
	EnvDesc_OTEL_HOST             = "OpenTelemetry 主機"

	EnvDesc_SUMMARIZER_PRESERVE_LAST       = "摘要器保留末段數"
	EnvDesc_SUMMARIZER_USE_QA              = "摘要器使用問答摘要"
	EnvDesc_SUMMARIZER_SUM_MSG_HUMAN_IN_QA = "摘要器在問答中包含使用者訊息"
	EnvDesc_SUMMARIZER_LAST_SEC_BYTES      = "摘要器末段位元組上限"
	EnvDesc_SUMMARIZER_MAX_BP_BYTES        = "摘要器系統提示詞位元組上限"
	EnvDesc_SUMMARIZER_MAX_QA_BYTES        = "摘要器問答位元組上限"
	EnvDesc_SUMMARIZER_MAX_QA_SECTIONS     = "摘要器問答區段上限"
	EnvDesc_SUMMARIZER_KEEP_QA_SECTIONS    = "摘要器保留問答區段"

	EnvDesc_ASSISTANT_SUMMARIZER_PRESERVE_LAST    = "助理摘要保留末段數"
	EnvDesc_ASSISTANT_SUMMARIZER_LAST_SEC_BYTES   = "助理摘要末段位元組上限"
	EnvDesc_ASSISTANT_SUMMARIZER_MAX_BP_BYTES     = "助理摘要系統提示詞位元組上限"
	EnvDesc_ASSISTANT_SUMMARIZER_MAX_QA_BYTES     = "助理摘要問答位元組上限"
	EnvDesc_ASSISTANT_SUMMARIZER_MAX_QA_SECTIONS  = "助理摘要問答區段上限"
	EnvDesc_ASSISTANT_SUMMARIZER_KEEP_QA_SECTIONS = "助理摘要保留問答區段"

	EnvDesc_EMBEDDING_PROVIDER        = "嵌入供應商"
	EnvDesc_EMBEDDING_URL             = "嵌入服務 URL"
	EnvDesc_EMBEDDING_KEY             = "嵌入 API 金鑰"
	EnvDesc_EMBEDDING_MODEL           = "嵌入模型"
	EnvDesc_EMBEDDING_BATCH_SIZE      = "嵌入批次大小"
	EnvDesc_EMBEDDING_STRIP_NEW_LINES = "嵌入時移除換行"
	EnvDesc_EMBEDDING_MAX_TEXT_BYTES  = "嵌入文字位元組上限"

	EnvDesc_ASK_USER = "人工介入"

	EnvDesc_ASSISTANT_USE_AGENTS = "為助理啟用多代理模式"

	EnvDesc_EXECUTION_MONITOR_ENABLED          = "啟用執行監控（beta）"
	EnvDesc_EXECUTION_MONITOR_SAME_TOOL_LIMIT  = "相同工具呼叫門檻"
	EnvDesc_EXECUTION_MONITOR_TOTAL_TOOL_LIMIT = "工具呼叫總數門檻"
	EnvDesc_MAX_GENERAL_AGENT_TOOL_CALLS       = "一般代理的最大工具呼叫數"
	EnvDesc_MAX_LIMITED_AGENT_TOOL_CALLS       = "受限代理的最大工具呼叫數"
	EnvDesc_AGENT_PLANNING_STEP_ENABLED        = "啟用任務規劃（beta）"

	EnvDesc_SCRAPER_PUBLIC_URL                    = "擷取器公開 URL"
	EnvDesc_SCRAPER_PRIVATE_URL                   = "擷取器私有 URL"
	EnvDesc_LOCAL_SCRAPER_USERNAME                = "本機擷取器使用者名稱"
	EnvDesc_LOCAL_SCRAPER_PASSWORD                = "本機擷取器密碼"
	EnvDesc_LOCAL_SCRAPER_MAX_CONCURRENT_SESSIONS = "擷取器最大並行工作階段數"

	EnvDesc_DUCKDUCKGO_ENABLED    = "DuckDuckGo 搜尋"
	EnvDesc_DUCKDUCKGO_REGION     = "DuckDuckGo 區域"
	EnvDesc_DUCKDUCKGO_SAFESEARCH = "DuckDuckGo 安全搜尋"
	EnvDesc_DUCKDUCKGO_TIME_RANGE = "DuckDuckGo 時間範圍"
	EnvDesc_SPLOITUS_ENABLED      = "Sploitus 搜尋"
	EnvDesc_PERPLEXITY_API_KEY    = "Perplexity API 金鑰"
	EnvDesc_TAVILY_API_KEY        = "Tavily API 金鑰"
	EnvDesc_FIRECRAWL_API_KEY     = "Firecrawl API 金鑰"
	EnvDesc_FIRECRAWL_API_URL     = "Firecrawl API URL"
	EnvDesc_TRAVERSAAL_API_KEY    = "Traversaal API 金鑰"
	EnvDesc_GOOGLE_API_KEY        = "Google 搜尋 API 金鑰"
	EnvDesc_GOOGLE_CX_KEY         = "Google 搜尋 CX 金鑰"
	EnvDesc_GOOGLE_LR_KEY         = "Google 搜尋 LR 金鑰"

	EnvDesc_DOCKER_INSIDE                    = "容器內的 Docker"
	EnvDesc_DOCKER_NET_ADMIN                 = "Docker 網路管理"
	EnvDesc_DOCKER_SOCKET                    = "Docker Socket 路徑"
	EnvDesc_DOCKER_NETWORK                   = "Docker 網路"
	EnvDesc_DOCKER_PUBLIC_IP                 = "Docker 對外 IP"
	EnvDesc_DOCKER_WORK_DIR                  = "Docker 工作目錄"
	EnvDesc_DOCKER_DEFAULT_IMAGE             = "Docker 預設映像"
	EnvDesc_DOCKER_DEFAULT_IMAGE_FOR_PENTEST = "Docker 滲透測試映像"
	EnvDesc_DOCKER_DEFAULT_IMAGE_FOR_TEST    = "Docker 沙箱自我測試映像"
	EnvDesc_DOCKER_IMAGE_SELECTION_MODE      = "Docker 映像選擇模式"
	EnvDesc_DOCKER_ALLOWED_IMAGES            = "Docker 允許的映像"
	EnvDesc_DOCKER_INSIDE_POLICY_TESTS       = "Docker 沙箱隔離自我測試"
	EnvDesc_DOCKER_HOST                      = "Docker 主機"
	EnvDesc_DOCKER_TLS_VERIFY                = "Docker TLS 驗證"
	EnvDesc_DOCKER_CERT_PATH                 = "Docker 憑證路徑"
	EnvDesc_DOCKER_INSIDE_HOST               = "Worker Docker Daemon 主機"
	EnvDesc_DOCKER_INSIDE_TLS_VERIFY         = "Worker Docker TLS 驗證"
	EnvDesc_DOCKER_INSIDE_CERT_PATH          = "Worker Docker 憑證路徑"

	EnvDesc_TENANT_ID                         = "PentAGI 租戶 ID"
	EnvDesc_LICENSE_KEY                       = "PentAGI 授權金鑰"
	EnvDesc_PPROF_ADDR                        = "PentAGI pprof 監聽位址"
	EnvDesc_PENTAGI_LISTEN_IP                 = "PentAGI 伺服器主機"
	EnvDesc_PENTAGI_LISTEN_PORT               = "PentAGI 伺服器連接埠"
	EnvDesc_PUBLIC_URL                        = "PentAGI 公開 URL"
	EnvDesc_CORS_ORIGINS                      = "PentAGI CORS 來源"
	EnvDesc_TRUSTED_PROXIES                   = "PentAGI 可信任代理"
	EnvDesc_COOKIE_SIGNING_SALT               = "PentAGI Cookie 簽章 Salt"
	EnvDesc_DATABASE_EXTENSIONS_SCHEMA        = "PostgreSQL 擴充結構"
	EnvDesc_DATABASE_SEARCH_PATH_VIA_OPTIONS  = "PostgreSQL 以選項設定搜尋路徑"
	EnvDesc_PROXY_URL                         = "HTTP/HTTPS 代理 URL"
	EnvDesc_HTTP_CLIENT_TIMEOUT               = "HTTP 用戶端逾時（秒）"
	EnvDesc_TERMINAL_TOOL_TIMEOUT             = "終端機工具逾時（秒）"
	EnvDesc_EXTERNAL_SSL_CA_PATH              = "自訂 CA 憑證路徑"
	EnvDesc_EXTERNAL_SSL_INSECURE             = "略過 SSL 驗證"
	EnvDesc_PENTAGI_SSL_DIR                   = "PentAGI SSL 目錄"
	EnvDesc_PENTAGI_DATA_DIR                  = "PentAGI 資料目錄"
	EnvDesc_PENTAGI_DOCKER_SOCKET             = "掛載 Docker Socket 路徑"
	EnvDesc_PENTAGI_DOCKER_CERT_PATH          = "掛載 Docker 憑證路徑"
	EnvDesc_PENTAGI_BEDROCK_CONFIG_PATH       = "AWS Bedrock 主機設定檔路徑"
	EnvDesc_PENTAGI_LLM_SERVER_CONFIG_PATH    = "自訂 LLM 主機設定檔路徑"
	EnvDesc_PENTAGI_OLLAMA_SERVER_CONFIG_PATH = "Ollama 主機設定檔路徑"

	EnvDesc_STATIC_DIR     = "前端靜態檔案目錄"
	EnvDesc_STATIC_URL     = "前端靜態檔案 URL"
	EnvDesc_SERVER_PORT    = "後端伺服器連接埠"
	EnvDesc_SERVER_HOST    = "後端伺服器主機"
	EnvDesc_SERVER_SSL_CRT = "後端伺服器 SSL 憑證路徑"
	EnvDesc_SERVER_SSL_KEY = "後端伺服器 SSL 金鑰路徑"
	EnvDesc_SERVER_USE_SSL = "後端伺服器使用 SSL"

	EnvDesc_PERPLEXITY_MODEL        = "Perplexity 模型"
	EnvDesc_PERPLEXITY_CONTEXT_SIZE = "Perplexity 上下文大小"
	EnvDesc_PERPLEXITY_TIMEOUT      = "Perplexity 逾時"

	EnvDesc_SEARXNG_URL        = "Searxng 搜尋 URL"
	EnvDesc_SEARXNG_CATEGORIES = "Searxng 搜尋分類"
	EnvDesc_SEARXNG_LANGUAGE   = "Searxng 搜尋語言"
	EnvDesc_SEARXNG_SAFESEARCH = "Searxng 安全搜尋"
	EnvDesc_SEARXNG_TIME_RANGE = "Searxng 時間範圍"
	EnvDesc_SEARXNG_TIMEOUT    = "Searxng 逾時"

	EnvDesc_OAUTH_GOOGLE_CLIENT_ID     = "OAuth Google 用戶端 ID"
	EnvDesc_OAUTH_GOOGLE_CLIENT_SECRET = "OAuth Google 用戶端密鑰"
	EnvDesc_OAUTH_GITHUB_CLIENT_ID     = "OAuth GitHub 用戶端 ID"
	EnvDesc_OAUTH_GITHUB_CLIENT_SECRET = "OAuth GitHub 用戶端密鑰"

	EnvDesc_LANGFUSE_EE_LICENSE_KEY   = "Langfuse 企業授權金鑰"
	EnvDesc_PENTAGI_POSTGRES_PASSWORD = "PentAGI PostgreSQL 密碼"

	EnvDesc_GRAPHITI_ENABLED                      = "啟用 Graphiti 整合"
	EnvDesc_GRAPHITI_URL                          = "Graphiti 伺服器 URL"
	EnvDesc_GRAPHITI_TIMEOUT                      = "Graphiti 請求逾時"
	EnvDesc_GRAPHITI_LLM_CLIENT_TYPE              = "Graphiti LLM 供應商預設集"
	EnvDesc_GRAPHITI_SEPARATE_EMBEDDING           = "Graphiti 獨立嵌入端點"
	EnvDesc_GRAPHITI_SEMAPHORE_LIMIT              = "Graphiti 協程上限"
	EnvDesc_GRAPHITI_LOG_LEVEL                    = "Graphiti 日誌等級"
	EnvDesc_GRAPHITI_SEARCH_SCOPE                 = "Graphiti 搜尋範圍"
	EnvDesc_GRAPHITI_INGEST_POLICY_RULES          = "Graphiti 匯入政策規則"
	EnvDesc_GRAPHITI_INGEST_POLICY_FIELD          = "Graphiti 匯入政策比對欄位"
	EnvDesc_GRAPHITI_INGEST_POLICY_DEFAULT_ACTION = "Graphiti 預設匯入動作"
	EnvDesc_GRAPHITI_INGEST_WORKER_COUNT          = "Graphiti 匯入 worker 數量"
	EnvDesc_GRAPHITI_INGEST_QUEUE_MAX_SIZE        = "Graphiti 匯入佇列上限"
	EnvDesc_GRAPHITI_TAXONOMY_LAYER_PROFILE       = "Graphiti 分類層設定檔"
	EnvDesc_GRAPHITI_CPUS                         = "Graphiti CPU 上限"
	EnvDesc_GRAPHITI_MEMORY                       = "Graphiti 記憶體上限"
	EnvDesc_NEO4J_USER                            = "Neo4j 使用者名稱"
	EnvDesc_NEO4J_DATABASE                        = "Neo4j 資料庫名稱"
	EnvDesc_NEO4J_PASSWORD                        = "Neo4j 資料庫密碼"
	EnvDesc_NEO4J_CPUS                            = "Neo4j CPU 上限"
	EnvDesc_NEO4J_MEMORY                          = "Neo4j 記憶體上限"
	EnvDesc_NEO4J_SHM_SIZE                        = "Neo4j 共享記憶體上限"
)

// dynamic, contextual sections used in processor operation forms
const (
	// section headers
	ProcessorSectionCurrentState = "目前狀態"
	ProcessorSectionPlanned      = "預計執行的動作"
	ProcessorSectionEffects      = "作用"

	// component labels
	ProcessorComponentPentagi       = "PentAGI"
	ProcessorComponentLangfuse      = "Langfuse"
	ProcessorComponentObservability = "可觀測性"

	ProcessorComponentWorkerImage           = "worker 映像"
	ProcessorComponentComposeStacks         = "compose 堆疊"
	ProcessorComponentDefaultFiles          = "預設檔案"
	ProcessorItemComposeFiles               = "compose 檔案"
	ProcessorItemComposeStacksImagesVolumes = "compose 堆疊、映像、資料卷"

	// common states
	ProcessorStateInstalled = "installed"
	ProcessorStateMissing   = "未安裝"
	ProcessorStateRunning   = "running"
	ProcessorStateStopped   = "stopped"
	ProcessorStateEmbedded  = "embedded"
	ProcessorStateExternal  = "external"
	ProcessorStateConnected = "connected"
	ProcessorStateDisabled  = "disabled"
	ProcessorStateUnknown   = "unknown"

	// planned action bullet prefixes
	PlannedWillStart    = "將啟動："
	PlannedWillStop     = "將停止："
	PlannedWillRestart  = "將重新啟動："
	PlannedWillUpdate   = "將更新："
	PlannedWillSkip     = "將略過："
	PlannedWillRemove   = "將移除："
	PlannedWillPurge    = "將清除："
	PlannedWillDownload = "將下載："
	PlannedWillRestore  = "將還原："

	// effect notes per operation (concise and practical)
	EffectsStart          = "PentAGI Web 介面即可使用，背景服務會依必要順序啟動。"
	EffectsStop           = "Web 介面將無法使用，進行中的流程會安全暫停；重新啟動 PentAGI 後流程會自動繼續，目前代理步驟的一小部分可能遺失。"
	EffectsRestart        = "服務會以乾淨狀態停止後重新啟動，預期中會有短暫中斷；之後流程會自動繼續。"
	EffectsUpdateAll      = "會拉取映像並視需要重建服務；外部或已停用的元件會略過。過程中預期會有短暫中斷。"
	EffectsDownloadWorker = "不會更動正在執行的 worker 容器。新流程會使用已下載的映像；若要讓既有流程改用新映像，請先結束該流程，再建立新任務或新助理。"
	EffectsUpdateWorker   = "拉取最新的 worker 映像。執行中的 worker 容器仍使用舊映像；新容器會使用更新後的映像。"
	EffectsFactoryReset   = "移除容器、資料卷與網路，並還原預設 .env 與內嵌檔案，回到乾淨的基準狀態。此操作無法復原。"
	EffectsRemove         = "停止並移除容器，但保留資料卷與映像，資料不會遺失。在你重新啟動之前，Web 介面將無法使用。"
	EffectsPurge          = "完整清除：容器、映像、資料卷與設定檔都會被刪除，且無法復原。"
	EffectsInstall        = "會建立必要檔案並啟動服務；偵測到外部元件時會略過。"
)
