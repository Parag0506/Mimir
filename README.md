# Mimir-AI: High-Performance, Cost-Efficient Coding Assistant

Mimir-AI is a self-hosted coding assistant that delivers **autocomplete**, **chat**, **refactoring**, and more—directly within VSCode. It harnesses powerful open-source models (Llama2, DeepSeek, CodeLlama) and optional proprietary models (GPT-4, Claude Sonnet), providing **low latency**, **scalable performance**, and **maximum flexibility**.

---

# Architecture

The system architecture includes:

- **VSCode Extension (TypeScript)**: Provides inline suggestions, chat, and refactoring commands.
- **API Gateway (Go)**: Receives requests from the extension, handles authentication, usage metering, and routing.
- **Model Router (Go)**: Determines which model instance to invoke, balances load, and leverages Redis for caching.
- **Model Manager (Python)**: Runs efficient inference (e.g., vLLM), includes 4-bit quantization (AWQ/GPTQ), supports fine-tuning.
- **Infrastructure**: GPU auto-scaling (Kubernetes + Karpenter), object storage for model weights, optional Spot Instances.
- **Monitoring**: Prometheus & Grafana for metrics and dashboards, ensuring real-time visibility.

---

# Repository Structure

```bash
mimir-ai
├── README.md                     # This file
├── .gitignore
├── LICENSE                       # MIT License file
├── CONTRIBUTING.md               # Contributing guidelines
├── docker-compose.yml            # Local dev environment
├── docker-compose.dev.yml        # Development specific compose
├── Makefile                      # Build and development commands
├── kubernetes/
│   ├── base/
│   │   ├── deployment.yaml
│   │   ├── service.yaml
│   │   └── kustomization.yaml
│   ├── overlays/
│   │   ├── development/
│   │   │   └── kustomization.yaml
│   │   └── production/
│   │       └── kustomization.yaml
│   ├── autoscaling.yaml
│   └── ingress.yaml
├── infra/
│   ├── terraform/
│   │   ├── main.tf
│   │   ├── variables.tf
│   │   ├── outputs.tf
│   │   ├── providers.tf
│   │   └── environments/
│   │       ├── dev/
│   │       └── prod/
│   └── scripts/
│       ├── setup.sh
│       ├── deploy.sh
│       └── backup.sh
├── extension/
│   ├── package.json
│   ├── tsconfig.json
│   ├── webpack.config.js
│   ├── .eslintrc.json
│   ├── jest.config.js
│   ├── src/
│   │   ├── extension.ts
│   │   ├── test/
│   │   ├── commands/
│   │   ├── services/
│   │   └── utils/
│   └── assets/
│       └── icon.png
├── services/
│   ├── api-gateway-go/
│   │   ├── Dockerfile
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── main.go
│   │   ├── internal/
│   │   │   ├── handlers/
│   │   │   ├── middleware/
│   │   │   └── config/
│   │   └── tests/
│   ├── model-router-go/
│   │   ├── Dockerfile
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── router.go
│   │   ├── internal/
│   │   │   ├── balancer/
│   │   │   ├── cache/
│   │   │   └── metrics/
│   │   └── tests/
│   └── model-manager-py/
│       ├── Dockerfile
│       ├── requirements.txt
│       ├── pyproject.toml
│       ├── manager.py
│       ├── tests/
│       └── mimir/
│           ├── __init__.py
│           ├── inference/
│           ├── optimization/
│           └── utils/
├── models/
│   ├── config/
│   │   ├── llama2.yaml
│   │   ├── deepseek.yaml
│   │   └── codellama.yaml
│   └── scripts/
│       ├── download.py
│       └── convert.py
├── cache/
│   └── redis/
│       ├── redis.conf
│       └── Dockerfile
├── monitoring/
│   ├── prometheus/
│   │   ├── prometheus.yml
│   │   └── alert.rules
│   └── grafana/
│       ├── dashboards/
│       └── datasources/
└── scripts/
    ├── build.sh
    ├── deploy.sh
    ├── test.sh
    └── lint.sh
```

---

# Getting Started

## Prerequisites
- **Docker & Docker Compose**: For local development.
- **Go 1.19+**: Required for the API Gateway and Model Router.
- **Python 3.9+**: For the Model Manager (vLLM or similar).
- **Node.js 16+**: To build the VSCode extension.
- **GPU**: Recommended for fast inference (CPU is possible but slower).

## Local Development

1. **Clone the Repository**:
```bash
git clone https://github.com/your-org/mimir-ai.git
cd mimir-ai
```
2. **Build & Start Services**:
```bash
docker-compose up --build
```
This brings up:
- Go-based **API Gateway**
- Go-based **Model Router**
- Python-based **Model Manager** (vLLM or equivalent)
- **Redis**, **Prometheus**, and **Grafana** (if configured)

3. **Install the VSCode Extension**:
```bash
cd extension
npm install
npm run build
# Install the resulting VSIX in VSCode
```

4. **Configure the Extension**:
- **API Endpoint**: `http://localhost:8080` or the port exposed by your gateway.
- **Model Selection**: e.g., `llama2-13b`, `deepseek-7b`.
- **Auth Token**: If tracking usage or using premium models.

## Usage
- **Autocomplete**: Type in VSCode to see inline suggestions.
- **Chat**: Open the Mimir-AI Chat command for Q&A, code explanations, or debugging.
- **Refactoring**: Highlight code and select a Mimir-AI refactoring action from the context menu.

---

# Model Serving & Optimization
- **vLLM / Text Generation Inference**: Provides highly efficient GPU-based model serving.
- **Quantization**: 4-bit AWQ or GPTQ to reduce VRAM usage while maintaining accuracy.
- **Batching**: Groups multiple user requests to improve throughput.
- **Fine-Tuning**: Enterprise tier allows custom training on private codebases.

---

# Deployment Options

1. **Self-Hosted Single GPU**:
   - Perfect for the **Free Tier** or individual use.
   - Deploy via `docker-compose` on your GPU machine.

2. **Kubernetes (Cloud or On-Prem)**:
   - Auto-scale GPU instances with Karpenter or similar tooling.
   - Compatible with AWS EKS, GCP GKE, Azure AKS, or self-managed clusters.

3. **Spot Instances**:
   - Leverage AWS EC2 Spot or equivalent for up to 90% cost savings.

4. **Hybrid**:
   - Proprietary models in the cloud, open-source models on-prem. Combines cost savings and privacy.

---

# Pricing & Tiers

- **Free ($0/month)**  
  - Self-host open-source models (Llama2, DeepSeek-7B).  
  - 100 user prompt credits/month for heavier tasks.

- **Developer ($10/month)**  
  - Premium open-source models (e.g., CodeLlama-34B).  
  - 1,000 user prompt credits/month.

- **Pro ($25/month)**  
  - Unlimited user prompt credits.  
  - Priority support via Slack/Email.

- **Teams ($20/user/month)**  
  - Shared prompt credits and usage analytics.  
  - Team-based administration and billing.

- **Enterprise (Custom)**  
  - Private fine-tuning, on-prem solutions, dedicated support.  
  - Custom SLAs and advanced security features.

---

# Monitoring & Logging

- **Prometheus**: Aggregates metrics like inference latency, GPU utilization, requests/sec.
- **Grafana**: Dashboards for performance and usage trends.
- **Logging**:  
  - Zero data retention by default (only usage metadata).  
  - Enterprise integrations (Elastic, Splunk) optional.

---

# Roadmap

1. **Phase 1 (2 months)**: Core extension, API Gateway (Go), Model Router (Go).
2. **Phase 2 (3 months)**: Integrate open-source models, Redis caching, performance tuning.
3. **Phase 3 (2 months)**: Self-hosting & cloud deployment (Docker/K8s).
4. **Phase 4 (1 month)**: Launch free & developer tiers with credit management.
5. **Phase 5 (2 months)**: Add pro, teams, enterprise tiers; fine-tuning & analytics capabilities.

---

# Contributing

We welcome any contributions! Please review our [CONTRIBUTING.md](./CONTRIBUTING.md) for guidelines on pull requests, coding standards, and testing.

---

# License

Mimir-AI is licensed under the [MIT License](./LICENSE). Some enterprise features or proprietary models may have additional terms.

---

# Contact & Support

- **Email**: support@mimir-ai.com
- **Community Chat**: (Slack/Discord link)
- **Issues**: (GitHub Issues link)

For enterprise inquiries or dedicated support, please contact our sales team.
