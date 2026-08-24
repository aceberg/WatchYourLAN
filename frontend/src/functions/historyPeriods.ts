export const HISTORY_DAY_START_HOUR = 7;
export const HISTORY_NIGHT_START_HOUR = 20;

export type HistoryPeriod = "day" | "night";

export type HistoryTimestamp = {
  year: number;
  month: number;
  day: number;
  hour: number;
  minute: number;
  second: number;
};

export function parseHistoryTimestamp(date: string): HistoryTimestamp | null {
  const match = /^(\d{4})-(\d{2})-(\d{2})[ T](\d{2}):(\d{2})(?::(\d{2}))?$/.exec(date);

  if (!match) {
    return null;
  }

  const [, year, month, day, hour, minute, second = "0"] = match;

  const timestamp = {
    year: Number(year),
    month: Number(month),
    day: Number(day),
    hour: Number(hour),
    minute: Number(minute),
    second: Number(second),
  };

  if (timestamp.month < 1 || timestamp.month > 12
    || timestamp.day < 1 || timestamp.day > 31
    || timestamp.hour < 0 || timestamp.hour > 23
    || timestamp.minute < 0 || timestamp.minute > 59
    || timestamp.second < 0 || timestamp.second > 59) {
    return null;
  }

  return timestamp;
}

export function getHistoryPeriod(date: string): HistoryPeriod | "" {
  const timestamp = parseHistoryTimestamp(date);

  if (!timestamp) {
    return "";
  }

  return timestamp.hour >= HISTORY_DAY_START_HOUR && timestamp.hour < HISTORY_NIGHT_START_HOUR
    ? "day"
    : "night";
}

export function historyPeriodLabel(period: HistoryPeriod | "") {
  if (period === "day") return "Day";
  if (period === "night") return "Night";
  return "";
}
