import { isDeviceTypeValue, type DeviceTypeValue } from "./deviceTypes";
import type { ActivityDeviceOption, ActivityStats, Conf, Host, HostEvent } from "./exports";

export const apiPath = '';
export type ActivityCategory = "all" | "connectivity" | "changes";
export type ActivityEventType = "discovered" | "online" | "offline" | "known" | "unknown" | "device-type-changed";

type ActivityQuery = {
  category?: ActivityCategory;
  eventTypes?: ActivityEventType[];
  offset?: number;
  mac?: string;
  macs?: string[];
};

const apiFetch = async (url: string, init?: RequestInit): Promise<Response> => {
  const response = await fetch(url, init);
  if (!response.ok) {
    const detail = await response.text();
    throw new Error(detail || response.statusText || "API request failed");
  }

  return response;
};

const apiJSON = async <T>(url: string, init?: RequestInit): Promise<T> => {
  return await (await apiFetch(url, init)).json();
};

export const apiGetAllHosts = async () => {
  const url = apiPath+'/api/all';
  const hosts = await apiJSON<Host[]>(url);

  return hosts;
};

export const apiGetConfig = async () => {

  const url = apiPath+'/api/config';
  const res = await apiJSON<Conf>(url);

  return res;
};

export const apiGetVersion = async () => {

  const url = apiPath+'/api/version';
  const res = await apiJSON<string>(url);

  return res;
};

export const apiGetActivity = async (limit = 20, query: ActivityQuery = {}): Promise<HostEvent[]> => {
  const params = new URLSearchParams({ limit: String(limit) });
  if (query.category !== undefined) {
    params.set("category", query.category);
  }
  query.eventTypes?.forEach((eventType) => {
    params.append("eventType", eventType);
  });
  if (query.offset !== undefined) {
    params.set("offset", String(query.offset));
  }
  if (query.mac) {
    params.append("mac", query.mac);
  }
  query.macs?.forEach((mac) => {
    if (mac) {
      params.append("mac", mac);
    }
  });

  const url = apiPath+'/api/activity?'+params.toString();
  const events = await apiJSON<HostEvent[]>(url);

  return events;
};

export const apiGetActivityStats = async (query: Pick<ActivityQuery, "mac" | "macs"> = {}): Promise<ActivityStats> => {
  const params = new URLSearchParams();
  if (query.mac) {
    params.append("mac", query.mac);
  }
  query.macs?.forEach((mac) => {
    if (mac) {
      params.append("mac", mac);
    }
  });

  const suffix = params.toString();
  const url = apiPath+'/api/activity/stats'+(suffix === "" ? "" : "?"+suffix);
  const stats = await apiJSON<ActivityStats>(url);

  return stats;
};

export const apiGetActivityDevices = async (): Promise<ActivityDeviceOption[]> => {
  const url = apiPath+'/api/activity/devices';
  const devices = await apiJSON<ActivityDeviceOption[]>(url);

  return devices;
};

export const apiGetHostActivity = async (id: number | string, limit = 10): Promise<HostEvent[]> => {
  const url = apiPath+'/api/host/'+id+'/activity?limit='+limit;
  const events = await apiJSON<HostEvent[]>(url);

  return events;
};

export const apiSetConfigColor = async (color: "dark" | "light"): Promise<Conf> => {
  const url = apiPath+'/api/config/color';
  return await apiJSON<Conf>(url, {
    method: 'POST',
    headers: { 'content-type': 'application/json' },
    body: JSON.stringify({ color }),
  });
};

export const apiSetRetention = async (presenceRetention: number, connectivityRetention: number): Promise<Conf> => {
  const url = apiPath+'/api/config/retention';
  return await apiJSON<Conf>(url, {
    method: 'POST',
    headers: { 'content-type': 'application/json' },
    body: JSON.stringify({ presenceRetention, connectivityRetention }),
  });
};

export const apiTestNotify = async () => {

  const url = apiPath+'/api/notify_test';
  await apiFetch(url);
};

export const apiEditHost = async (id:number, name:string, known:string) => {

  const url = apiPath+'/api/edit/'+id+'/'+name+'/'+known;
  const res = await apiJSON<string>(url);

  return res;
};

export const apiSetDeviceType = async (id: number, deviceType: DeviceTypeValue): Promise<Host> => {
  if (!isDeviceTypeValue(deviceType)) {
    throw new Error("invalid device type");
  }

  const url = apiPath+'/api/host/'+id+'/type';
  return await apiJSON<Host>(url, {
    method: 'PATCH',
    headers: { 'content-type': 'application/json' },
    body: JSON.stringify({ deviceType }),
  });
};

export const apiGetHost = async (id:string) => {

  const url = apiPath+'/api/host/'+id;
  const res = await apiJSON<Host>(url);

  return res;
};

export const apiDelHost = async (id:number) => {

  const url = apiPath+'/api/host/del/'+id;
  const res = await apiJSON<string>(url);

  return res;
};

export const apiPortScan = async (ip:string, port:number) => {

  const url = apiPath+'/api/port/'+ip+'/'+port;
  const res = await apiJSON<boolean>(url);

  return res;
};

export const apiGetHistory = async (mac:string) => {
  const url = apiPath+'/api/history/'+mac+'/?num=210';
  const hosts = await apiJSON<Host[]>(url);

  return hosts;
};

export const apiGetHistoryByDate = async (mac:string, date: string) => {
  const url = apiPath+'/api/history/'+mac+'/'+date;
  const hosts = await apiJSON<Host[]>(url);

  return hosts;
};

export const apiWOL = async (mac:string) => {

  const url = apiPath+'/api/wol/'+mac;
  const res = await apiJSON<boolean>(url);

  return res;
};
