# Gira

A blazingly fast, read-only Jira CLI client written in Go. Gira provides a beautiful terminal user interface (TUI) for viewing and navigating your Jira active sprints - much faster than the web interface.

> **Note**: This is currently a demo project. The application is read-only; users cannot update or edit sprint issues.


https://github.com/user-attachments/assets/343b4fde-4831-43e9-b9f4-c47021dfba8f


## Features

- **Fast Performance**: Terminal-based interface that's significantly faster than Jira's web UI
- **Active Sprint Viewer**: Quickly view and navigate active sprint issues
- **Project & Board Navigation**: Browse through Jira projects and boards
- **Task Board View**: Organized view of sprint tasks by status
- **Task Details**: View detailed information about individual tasks
- **Read-Only**: Safe, non-destructive viewing of your Jira data
- **Beautiful TUI**: Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and styled with [Lipgloss](https://github.com/charmbracelet/lipgloss)

## Prerequisites

- Go 1.24.6 or higher
- A Jira account with API access
- Jira API token

## Installation

1. Clone the repository:
```bash
git clone https://github.com/mabd-dev/gira.git
cd gira
```

2. Install dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build -o gira
```

## Configuration

Gira uses a TOML configuration file for authentication with Jira. On first run, the configuration file is automatically created at `~/.config/gira/credentials.toml` with secure permissions.

1. Run the application for the first time to create the config file:
```bash
./gira
```

The application will create `~/.config/gira/credentials.toml` and prompt you to edit it.

2. Edit the configuration file with your Jira credentials:
```bash
nano ~/.config/gira/credentials.toml
```

3. Update the following fields:
```toml
[general]
debug = false

[credentials]
email = "your.email@example.com"
secret = "your_jira_api_token"
domain = "your_jira_domain"
```

### Getting Your Jira API Token

1. Log in to your Jira account
2. Go to [Atlassian Account Settings](https://id.atlassian.com/manage-profile/security/api-tokens)
3. Click "Create API token"
4. Give it a label and click "Create"
5. Copy the generated token to your `credentials.toml` file as the `secret` value

### Configuration Fields

- **general.debug**: Enable debug mode to use mock API client (default: `false`)
- **credentials.email**: Your Jira account email address
- **credentials.secret**: Your Jira API token (see above)
- **credentials.domain**: Your Jira domain (e.g., if your Jira URL is `yourcompany.atlassian.net`, use `yourcompany`)

## Usage

Run the application:

```bash
./gira
```

Or run directly with Go:

```bash
go run main.go
```

### Navigation

The TUI provides an interactive interface for navigating your Jira workspace:

1. **Projects View**: Browse available Jira projects
2. **Boards View**: Select from available boards in a project
3. **Sprint View**: View the active sprint with all issues
4. **Task Board**: See tasks organized by status columns
5. **Task Details**: View detailed information about individual tasks

## Project Structure

```
gira/
├── api/                   # Jira API client implementation
│   ├── testdata/          # mock data for testing
│   ├── client.go          # Main client interface
│   ├── clientReal.go      # Real API client
│   ├── clientMock.go      # Mock client for testing
│   └── types.go           # API response types
├── config/                # Configuration management
│   └── credentials.go     # Environment variable loading
├── models/                # Domain models
│   ├── types.go           # Core domain types
│   ├── formatter.go       # Data formatters
│   └── taskStatus.go      # Task status definitions
├── internal/
│   ├── logger/            # Logging functionality
│   ├── theme/             # UI theming
│   └── ui/                # TUI components
│       ├── projects/      # Projects view
│       ├── boards/        # Boards view
│       └── sprint/        # Sprint views
│           ├── tasksboard/    # Task board view
│           └── taskdetails/   # Task details view
└── main.go                # Application entry point
```

## Technologies Used

- **[Go](https://golang.org/)**: Programming language
- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)**: TUI framework
- **[Bubbles](https://github.com/charmbracelet/bubbles)**: TUI components
- **[Lipgloss](https://github.com/charmbracelet/lipgloss)**: Terminal styling
- **[go-toml/v2](https://github.com/pelletier/go-toml)**: TOML configuration management
- **Jira REST API**: Data source

## Development

### Logs

Logs are stored in `~/.config/gira/logs/` for debugging and troubleshooting.

## Architecture

Gira follows a clean architecture pattern:

- **API Layer**: Handles communication with Jira REST API
- **Models Layer**: Domain models independent of API responses
- **UI Layer**: Terminal user interface components using Bubble Tea
- **Config Layer**: Configuration and credentials management

The application uses the Bubble Tea framework's Model-View-Update (MVU) pattern for state management and rendering.

## Limitations

As a demo project, Gira currently has the following limitations:

- **Read-only**: No ability to create, update, or delete issues
- **Active Sprint Only**: Only displays the currently active sprint
- **Basic Views**: Limited to projects, boards, and sprint issues

## Future Enhancements

Potential features for future development:

- Issue editing and creation
- Multiple sprint viewing
- Custom filters and search
- Offline mode with caching
- Custom themes and color schemes
- Comment viewing

## Contributing

This is a demo project, but contributions are welcome! Feel free to:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

## Acknowledgments

- Built with [Charm](https://charm.sh/) libraries
- Inspired by the need for a faster Jira experience

