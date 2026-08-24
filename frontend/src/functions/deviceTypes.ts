import type { Host } from "./exports";

export type DeviceTypeValue =
  | ""
  | "router"
  | "switch"
  | "access-point"
  | "firewall"
  | "server"
  | "nas"
  | "desktop"
  | "laptop"
  | "phone"
  | "tablet"
  | "tv"
  | "printer"
  | "camera"
  | "iot"
  | "virtual-machine"
  | "container"
  | "game-console"
  | "other";

export type DeviceTypeTone = "unassigned" | "infrastructure" | "default";

export type DeviceTypeOption = {
  value: DeviceTypeValue;
  label: string;
  icon: string;
  tone: DeviceTypeTone;
};

export type DeviceTypeFilterOption = {
  value: string;
  label: string;
};

export const deviceTypeNotSetFilterValue = "not-set";
const legacyDeviceTypeNotSetFilterValue = "__not-set";

export const deviceTypes: readonly DeviceTypeOption[] = [
  { value: "", label: "Not set", icon: "bi-question-circle", tone: "unassigned" },
  { value: "router", label: "Router", icon: "bi-router-fill", tone: "infrastructure" },
  { value: "switch", label: "Switch", icon: "bi-diagram-3-fill", tone: "infrastructure" },
  { value: "access-point", label: "Access Point", icon: "bi-wifi", tone: "infrastructure" },
  { value: "firewall", label: "Firewall", icon: "bi-shield-fill", tone: "infrastructure" },
  { value: "server", label: "Server", icon: "bi-server", tone: "default" },
  { value: "nas", label: "NAS", icon: "bi-hdd-stack-fill", tone: "default" },
  { value: "desktop", label: "Desktop", icon: "bi-pc-display-horizontal", tone: "default" },
  { value: "laptop", label: "Laptop", icon: "bi-laptop", tone: "default" },
  { value: "phone", label: "Phone", icon: "bi-phone", tone: "default" },
  { value: "tablet", label: "Tablet", icon: "bi-tablet", tone: "default" },
  { value: "tv", label: "TV / Media Device", icon: "bi-tv", tone: "default" },
  { value: "printer", label: "Printer", icon: "bi-printer", tone: "default" },
  { value: "camera", label: "Camera", icon: "bi-camera-video-fill", tone: "default" },
  { value: "iot", label: "IoT Device", icon: "bi-cpu-fill", tone: "default" },
  { value: "virtual-machine", label: "Virtual Machine", icon: "bi-window-stack", tone: "default" },
  { value: "container", label: "Container", icon: "bi-box-seam-fill", tone: "default" },
  { value: "game-console", label: "Game Console", icon: "bi-controller", tone: "default" },
  { value: "other", label: "Other", icon: "bi-three-dots", tone: "default" },
];

const deviceTypeByValue = new Map(deviceTypes.map((type) => [type.value, type]));

export function normalizeDeviceType(value: string | null | undefined): DeviceTypeValue {
  return deviceTypeByValue.has(value as DeviceTypeValue) ? value as DeviceTypeValue : "";
}

export function isDeviceTypeValue(value: string): value is DeviceTypeValue {
  return deviceTypeByValue.has(value as DeviceTypeValue);
}

export function getDeviceTypeOption(value: string | null | undefined): DeviceTypeOption {
  return deviceTypeByValue.get(normalizeDeviceType(value)) ?? deviceTypes[0];
}

export function normalizeDeviceTypeFilter(value: unknown): string {
  if (value === deviceTypeNotSetFilterValue || value === legacyDeviceTypeNotSetFilterValue) {
    return deviceTypeNotSetFilterValue;
  }
  if (typeof value !== "string") {
    return "";
  }

  const normalized = normalizeDeviceType(value);
  return normalized === "" ? "" : normalized;
}

export function deviceTypeFilterMatches(host: Host, filterValue: string): boolean {
  const normalizedFilter = normalizeDeviceTypeFilter(filterValue);
  if (normalizedFilter === "") {
    return true;
  }

  const hostDeviceType = normalizeDeviceType(host.DeviceType);
  if (normalizedFilter === deviceTypeNotSetFilterValue) {
    return hostDeviceType === "";
  }

  return hostDeviceType === normalizedFilter;
}

export function deviceTypeFilterLabel(value: string): string {
  const normalized = normalizeDeviceTypeFilter(value);
  if (normalized === "") {
    return "All types";
  }
  if (normalized === deviceTypeNotSetFilterValue) {
    return deviceTypes[0].label;
  }

  return getDeviceTypeOption(normalized).label;
}

export function deviceTypeFilterOptions(hosts: Host[], activeValue = ""): DeviceTypeFilterOption[] {
  const presentTypes = new Set(hosts.map((host) => normalizeDeviceType(host.DeviceType)));
  const normalizedActive = normalizeDeviceTypeFilter(activeValue);
  const options: DeviceTypeFilterOption[] = [{ value: "", label: "All types" }];

  for (const type of deviceTypes) {
    const value = type.value === "" ? deviceTypeNotSetFilterValue : type.value;
    const isPresent = type.value === "" ? presentTypes.has("") : presentTypes.has(type.value);

    if (isPresent || normalizedActive === value) {
      options.push({ value, label: type.label });
    }
  }

  return options;
}
