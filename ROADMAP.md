# Sonde Roadmap:
- [ ] Phase 1: Initial CLI
  - [X] Pipe in logs from your terminal
  - [ ] Summarize command:
    - [ ] sonde summarize (summarize the logs for you... in case you don't like reading)
      - [ ] summarize
      - [ ] s (shorthand)
  - [ ] Retrieve logs from previously ran commands
    - [X] terminal history to rerun commands
      - [X] fish
      - [ ] bash
      - [ ] zsh
    - [ ] tmux and others to get logs without rerunning commands
  - [ ] Initial connections to local LLM

- [ ] Initial Release!

- [ ] Phase 2: Container log retrieval
  - [ ] Retrieve logs from docker containers
  - [ ] Retrieve logs from podman containers
  - [ ] Retrieve logs from kubernetes pods

- [ ] Phase 3: adding commands
  - [ ] sonde debug (Help you debug the logs, give actionable insights)
    - [ ] debug
    - [ ] d (shorthand)
  - [ ] sonde cleanup (Suggestions for you to cleanup the logs, remove sensitive info, duplicates, etc.)
    - [ ] cleanup
    - [ ] c (shorthand)
  - [ ] sonde metrics (Show you metrics about the logs, like how many errors, warnings, etc.)
    - [ ] metrics
    - [ ] m (shorthand)
  - [ ] sonde explain (Explain the logs for you, in case you don't understand what's going on)
    - [ ] explain
    - [ ] e (shorthand)
  - [ ] sonde audit (Looks at logs and look for active/escalating anomalies. Looking for potential security issues)
    - [ ] audit
    - [ ] a (shorthand)
  - [ ] sonde translate (Translate the logs into a file format (e.g. JSON, CSV)
    - [ ] translate
    - [ ] t (shorthand)
  - [ ] sonde inspect (Inspect the logs against standard best practices)
    - [ ] inspect
    - [ ] i (shorthand)

- [ ] Phase 4: Add flags for commands
  - [ ] Output to a file (Overwrites or creates a new file)
    - [ ] --output
    - [ ] -o (shorthand)
  - [ ] Append to a file
    - [ ] --append
    - [ ] -a (shorthand)
  - [ ] Model (change LLM model on the fly without updating the config)
    - [ ] --model
    - [ ] -m (shorthand)
  - [ ] Temperature (Set the temperature for the LLM model)
    - [ ] --temperature
    - [ ] -t (shorthand)
  - [ ] Provider (Specify the provider for the LLM model)
    - [ ] --provider
    - [ ] -p (shorthand)
  - [ ] Api Key (Override the API key for the LLM model)
    - [ ] --api-key
    - [ ] -k (shorthand)
  - [ ] Config (Override the config file for the LLM model)
    - [ ] --config
    - [ ] -c (shorthand)
  - [ ] Raw (Bypass sonde filtering, directly take all the logs as-is)
    - [ ] --raw
    - [ ] -r (shorthand)
  - [ ] Filter (Apply a filter to the logs)
    - [ ] --filter
    - [ ] -f (shorthand)
  - [ ] Lines (Number of lines from the bottom of the logs to send)
    - [ ] --lines
    - [ ] -l (shorthand)
  - [ ] Silent (Suppress output to stdout)
    - [ ] --silent
    - [ ] -s (shorthand)

- [ ] Phase 5: Rolling out more LOCAL LLMs
  - [ ] LocalAI
  - [ ] Open to ideas lol

- [ ] Phase 6: Development work for "dredge"
  - [ ] Figure out Zero Trust store
  - [ ] sanitize with CLI
  - [ ] restore with CLI

- [ ] Phase 7: Rollout cloud LLMs
  - [ ] Amazon Bedrock
  - [ ] Anthropic
  - [ ] OpenAI
  - [ ] Gemini/Google

- [ ] Phase 8: Caching and Context (For CLI and web UI)
  - [ ] Research caching and context strategies
  - [ ] Implement caching and context strategies
  - [ ] Research a way to link to frontend of chosen LLM provider
  - [ ] Create REST API
    - [ ] API to receive context from CLI
    - [ ] Able to save context in DB
    - [ ] Able to connect to AI from REST API
    - [ ] Different login methods
      - [ ] OIDC
      - [ ] Basic Auth
      - [ ] Token-based
  - [ ] Create UI
    - [ ] Able to view context
    - [ ] Able to keep conversation going from CLI tool
    - [ ] Login
      - [ ] OIDC
      - [ ] Basic Auth
      - [ ] Token-based
    - [ ] Logout

- [ ] Phase 9: Rollout Daemon
  - [ ] Create Daemon that can be deployed as a service
  - [ ] Daemon collects terminal output to create better context to send to the AI model

- [ ] Phase 10: RAG
  - [ ] Create RAG to ingest user's current codebase for context
  - [ ] Should also go through "dredge" to remove sensitive information before sending to the AI model

Other things to consider
- Auto-completion
