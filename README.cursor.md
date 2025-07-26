# README.cursor.md

## 📦 Repo: lawnibly-scheduler-service

### 🎯 Purpose
Microservice in Go responsible for job scheduling, frequency tracking, and reminders.

### 🧠 Context
- Interfaces with backend API for job data
- Future role: async job assignment, reminder queues, provider availability checks

### 🔧 Current State
- Initial setup with Go (ready to be extended)
- Will manage weekly/bi-weekly recurrence jobs
- Placeholder for future cron or queue workers

### 🔜 Next Features
- Connect to PostgreSQL or message bus (via backend)
- Auto-create jobs based on booking recurrence

### 🧭 Connected Docs
- [docs/development/scheduler.md](../../docs/development/scheduler.md)
