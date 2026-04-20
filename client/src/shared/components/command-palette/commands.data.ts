export type CommandGroup = "Navigation" | "Infrastructure" | "Actions" | "Settings";

export interface Cmd {
  id: string;
  group: CommandGroup;
  label: string;
  hint?: string;
  shortcut?: string;
  to?: string;
}

export const commands: Cmd[] = [
  { id: "nav-overview", group: "Navigation", label: "Overview", hint: "Dashboard principal", to: "/" },
  { id: "nav-logs", group: "Navigation", label: "Logs", hint: "Stream temps réel", to: "/logs" },
  { id: "nav-alerts", group: "Navigation", label: "Alerts", hint: "Alertes et règles", to: "/alerts" },
  { id: "nav-settings", group: "Navigation", label: "Settings", hint: "Configuration", to: "/settings" },

  { id: "infra-routers", group: "Infrastructure", label: "Routers" },
  { id: "infra-services", group: "Infrastructure", label: "Services" },
  { id: "infra-middlewares", group: "Infrastructure", label: "Middlewares" },
  { id: "infra-certs", group: "Infrastructure", label: "Certificates" },

  { id: "act-refresh", group: "Actions", label: "Actualiser les panneaux", shortcut: "R" },
  { id: "act-export", group: "Actions", label: "Exporter le dashboard", shortcut: "⌘ E" },
  { id: "act-panel", group: "Actions", label: "Ajouter un panneau", shortcut: "⌘ N" },
  { id: "act-theme", group: "Actions", label: "Basculer le thème" },

  { id: "set-switch-instance", group: "Settings", label: "Changer d'instance" },
  { id: "set-shortcuts", group: "Settings", label: "Raccourcis clavier", shortcut: "?" },
];
