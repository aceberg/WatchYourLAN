import { getDeviceTypeOption } from "./deviceTypes";
import type { HostEvent } from "./exports";

export type ActivityTone = "discovered" | "online" | "offline" | "known" | "unknown" | "type";

export function activityIcon(eventType: string): string {
  switch (eventType) {
    case "online":
      return "bi-check-circle-fill";
    case "offline":
      return "bi-x-circle-fill";
    case "known":
      return "bi-bookmark-check-fill";
    case "unknown":
      return "bi-question-circle-fill";
    case "device-type-changed":
      return "bi-tag-fill";
    case "discovered":
    default:
      return "bi-plus-circle-fill";
  }
}

export function activityTone(eventType: string): ActivityTone {
  switch (eventType) {
    case "online":
      return "online";
    case "offline":
      return "offline";
    case "known":
      return "known";
    case "unknown":
      return "unknown";
    case "device-type-changed":
      return "type";
    case "discovered":
    default:
      return "discovered";
  }
}

export function activityDescription(event: HostEvent): string {
  switch (event.EventType) {
    case "online":
      return "Came online";
    case "offline":
      return "Went offline";
    case "known":
      return "Marked as known";
    case "unknown":
      return "Marked as unknown";
    case "device-type-changed": {
      return "Type changed to " + activityDeviceTypeLabel(event.NewValue);
    }
    case "discovered":
    default:
      return "New device detected";
  }
}

export function activityDetails(event: HostEvent): string {
  switch (event.EventType) {
    case "device-type-changed":
      return "Type changed: " + activityDeviceTypeLabel(event.OldValue) + " \u2192 " + activityDeviceTypeLabel(event.NewValue);
    case "discovered":
    case "known":
    case "unknown":
      return compactNetworkDetail(event);
    case "online":
    case "offline":
    default:
      return "";
  }
}

export function activityHostName(event: HostEvent): string {
  return event.Name.trim() || event.Mac || "Unknown device";
}

export function activityDeviceIcon(event: HostEvent): string {
  return getDeviceTypeOption(event.DeviceType).icon;
}

export function activityDeviceTypeLabel(value: string | null | undefined): string {
  return getDeviceTypeOption(value).label;
}

export function activityCategoryLabel(eventType: string): string {
  switch (eventType) {
    case "online":
    case "offline":
      return "Connectivity";
    default:
      return "Device changes";
  }
}

export function activityEventLabel(eventType: string): string {
  switch (eventType) {
    case "online":
      return "Online";
    case "offline":
      return "Offline";
    case "known":
      return "Marked known";
    case "unknown":
      return "Marked unknown";
    case "device-type-changed":
      return "Device type changed";
    case "discovered":
    default:
      return "New device detected";
  }
}

export function activityDayLabel(value: string): string {
  const match = value.match(/^(\d{4})-(\d{2})-(\d{2})/);
  if (!match) {
    return "Unknown day";
  }

  const [, year, month, day] = match;
  const date = new Date(Number(year), Number(month) - 1, Number(day));
  return date.toLocaleDateString("en", {
    day: "2-digit",
    month: "short",
    year: "numeric",
  });
}

export function relativeActivityTime(value: string): string {
  const date = parseActivityDate(value);
  if (date === null) {
    return "";
  }

  const seconds = Math.max(0, Math.floor((Date.now() - date.getTime()) / 1000));
  if (seconds < 60) {
    return "now";
  }

  const minutes = Math.floor(seconds / 60);
  if (minutes < 60) {
    return minutes + " min ago";
  }

  const hours = Math.floor(minutes / 60);
  if (hours < 24) {
    return hours + " h ago";
  }

  if (hours < 48) {
    return "yesterday";
  }

  return Math.floor(hours / 24) + " d ago";
}

function compactNetworkDetail(event: HostEvent): string {
  if (event.IP && event.Iface) {
    return event.IP + " / " + event.Iface;
  }

  return event.IP || event.Iface || "";
}

function parseActivityDate(value: string): Date | null {
  const match = value.match(/^(\d{4})-(\d{2})-(\d{2})[ T](\d{2}):(\d{2}):(\d{2})$/);
  if (!match) {
    return null;
  }

  const [, year, month, day, hour, minute, second] = match;
  return new Date(
    Number(year),
    Number(month) - 1,
    Number(day),
    Number(hour),
    Number(minute),
    Number(second),
  );
}
