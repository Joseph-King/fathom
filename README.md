# Sonde
Sonde currently is a CLI based tool writtin in Go. It's main purpose is to send logs from your terminal and docker/podman/kube containers, throw them into an LLM, and give you a human readable insights. It's target environment is for local development and debugging, but plans are in the works to leverage cloud LLMs leading to a fully distributed logging and debugging system.

There are plans to flesh this out into a ecosystem of tools working together to provide a comprehensive logging and debugging experience. Which includes:
  - CLI tool for sending logs
  - Web UI for visualizing logs
  - API for programmatic access to logs
  - Daemon for streamlining log collection
  - RAG to get context into codebase being debugged
  - Separate tool codename "dredge" for sanitizing logs via a zero trust store, so that it is able to restore context.

CURRENTLY IN ACTIVE DEVELOPMENT

See ROADMAP.md for the current roadmap.
