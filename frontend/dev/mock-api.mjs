import { createReadStream, existsSync } from 'node:fs';
import { createServer } from 'node:http';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const host = '127.0.0.1';
const port = 8840;
const now = '2026-08-23 10:15:00';
let nextActivityId = 1;
const deviceTypes = new Set([
  '',
  'router',
  'switch',
  'access-point',
  'firewall',
  'server',
  'nas',
  'desktop',
  'laptop',
  'phone',
  'tablet',
  'tv',
  'printer',
  'camera',
  'iot',
  'virtual-machine',
  'container',
  'game-console',
  'other',
]);
const connectivityEvents = new Set(['online', 'offline']);
const changeEvents = new Set(['discovered', 'known', 'unknown', 'device-type-changed']);
const validActivityEvents = new Set([...connectivityEvents, ...changeEvents]);

const fakeHosts = [
  {
    ID: 1,
    Name: 'router',
    DNS: 'router.local',
    Iface: 'eth0',
    IP: '192.168.1.1',
    Mac: 'AA:BB:CC:00:00:01',
    Hw: 'Example Networks Gateway',
    Date: now,
    Known: 1,
    Now: 1,
    DeviceType: 'router',
  },
  {
    ID: 2,
    Name: 'NAS',
    DNS: 'nas.local',
    Iface: 'eth0',
    IP: '192.168.1.20',
    Mac: 'AA:BB:CC:00:00:20',
    Hw: 'Storage Appliance',
    Date: now,
    Known: 1,
    Now: 1,
    DeviceType: 'nas',
  },
  {
    ID: 3,
    Name: 'desktop',
    DNS: 'desktop.local',
    Iface: 'eth0',
    IP: '192.168.1.42',
    Mac: 'AA:BB:CC:00:00:42',
    Hw: 'Workstation NIC',
    Date: now,
    Known: 1,
    Now: 1,
    DeviceType: 'desktop',
  },
  {
    ID: 4,
    Name: 'phone',
    DNS: 'phone.local',
    Iface: 'wifi0',
    IP: '192.168.1.83',
    Mac: 'AA:BB:CC:00:00:83',
    Hw: 'Mobile Device',
    Date: now,
    Known: 0,
    Now: 1,
    DeviceType: 'phone',
  },
  {
    ID: 5,
    Name: 'offline device',
    DNS: '',
    Iface: 'eth0',
    IP: '192.168.1.120',
    Mac: 'AA:BB:CC:00:01:20',
    Hw: 'Legacy Device',
    Date: '2025-12-30 18:42:09',
    Known: 1,
    Now: 0,
    DeviceType: '',
  },
];

const config = {
  Host: host,
  Port: String(port),
  Theme: 'sand',
  Color: 'dark',
  DirPath: './dev/.mock-data',
  ConfPath: './dev/.mock-data/config_v2.yaml',
  DBPath: './dev/.mock-data/scan.db',
  NodePath: '',
  LogLevel: 'info',
  Ifaces: 'eth0 wifi0',
  ArpArgs: '',
  ArpStrs: [],
  Timeout: 600,
  TrimHist: 48,
  ConnectivityRetention: 168,
  ShoutURL: '',
  ShoutURLConfigured: false,
  Version: 'dev-mock',
  UseDB: 'sqlite',
  PGConnect: '',
  PGConnectConfigured: false,
  InfluxEnable: false,
  InfluxAddr: '',
  InfluxToken: '',
  InfluxTokenConfigured: false,
  InfluxOrg: '',
  InfluxBucket: '',
  InfluxSkipTLS: false,
  PrometheusEnable: false,
};

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const frontendPublicPath = path.resolve(__dirname, '../public/fs/public');
const backendPublicPath = path.resolve(__dirname, '../../backend/internal/web/public');
const localPublicAssets = new Set([
  'favicon.png',
  'lanventory.ico',
  'lanventory-16x16.png',
  'lanventory-32x32.png',
  'lanventory-48x48.png',
  'lanventory-64x64.png',
  'lanventory-128x128.png',
  'lanventory-180x180.png',
  'lanventory-192x192.png',
  'lanventory-256x256.png',
  'lanventory-512x512.png',
  'lanventory-navbar.png',
]);
const activityEvents = [];

function sendJSON(res, value, statusCode = 200) {
  res.writeHead(statusCode, {
    'content-type': 'application/json; charset=utf-8',
    'cache-control': 'no-store',
  });
  res.end(JSON.stringify(value));
}

function sendText(res, value, statusCode = 200) {
  res.writeHead(statusCode, {
    'content-type': 'text/plain; charset=utf-8',
    'cache-control': 'no-store',
  });
  res.end(value);
}

function getLocalPublicAsset(pathname) {
  const match = /^\/fs\/public\/([^/]+)$/.exec(pathname);
  if (!match || !localPublicAssets.has(match[1])) {
    return '';
  }

  for (const root of [frontendPublicPath, backendPublicPath]) {
    const assetPath = path.join(root, match[1]);
    if (existsSync(assetPath)) {
      return assetPath;
    }
  }

  return '';
}

function getAssetContentType(assetPath) {
  return assetPath.endsWith('.ico') ? 'image/x-icon' : 'image/png';
}

function readBody(req) {
  return new Promise((resolve) => {
    let body = '';
    req.on('data', (chunk) => {
      body += chunk;
    });
    req.on('end', () => resolve(body));
    req.on('error', () => resolve(body));
  });
}

function isColorMode(color) {
  return color === 'dark' || color === 'light';
}

function isDeviceType(deviceType) {
  return typeof deviceType === 'string' && deviceTypes.has(deviceType);
}

function isPositiveIntegerValue(value) {
  const parsed = Number(value);
  return Number.isInteger(parsed) && parsed > 0;
}

function parseRequestBody(body) {
  try {
    const parsed = JSON.parse(body);
    return parsed && typeof parsed === 'object' ? parsed : {};
  } catch {
    return Object.fromEntries(new URLSearchParams(body));
  }
}

function publicConfig() {
  const { ConfPath, DBPath, Version, ...publicFields } = config;

  return {
    ...publicFields,
    ShoutURL: '',
    ShoutURLConfigured: config.ShoutURL !== '',
    PGConnect: '',
    PGConnectConfigured: config.PGConnect !== '',
    InfluxToken: '',
    InfluxTokenConfigured: config.InfluxToken !== '',
  };
}

function isTruthyFormValue(value) {
  return ['1', 'on', 'true', 'yes'].includes(String(value ?? '').trim().toLowerCase());
}

function applySecretUpdate(currentValue, submittedValue, clearValue) {
  if (isTruthyFormValue(clearValue)) {
    return '';
  }
  if (typeof submittedValue === 'string' && submittedValue !== '') {
    return submittedValue;
  }

  return currentValue;
}

function applyBasicConfigForm(body) {
  const form = parseRequestBody(body);

  if (typeof form.host === 'string') {
    config.Host = form.host;
  }
  if (typeof form.port === 'string') {
    config.Port = form.port;
  }
  if (typeof form.theme === 'string' && /^[a-z0-9-]+$/.test(form.theme)) {
    config.Theme = form.theme;
  }
  if (isColorMode(form.color)) {
    config.Color = form.color;
  }
  if (typeof form.node === 'string') {
    config.NodePath = form.node;
  }
  config.ShoutURL = applySecretUpdate(config.ShoutURL, form.shout, form.clear_shout);
}

function applySettingsConfigForm(body) {
  const form = parseRequestBody(body);

  if (typeof form.ifaces === 'string') {
    config.Ifaces = form.ifaces;
  }
  if (typeof form.arpargs === 'string') {
    config.ArpArgs = form.arpargs;
  }
  if (typeof form.log === 'string') {
    config.LogLevel = form.log;
  }
  if (isPositiveIntegerValue(form.timeout)) {
    config.Timeout = Number(form.timeout);
  }
  if (isPositiveIntegerValue(form.trim)) {
    config.TrimHist = Number(form.trim);
  }
  if (isPositiveIntegerValue(form.connectivity_retention)) {
    config.ConnectivityRetention = Number(form.connectivity_retention);
  }
  if (typeof form.usedb === 'string') {
    config.UseDB = form.usedb;
  }
  config.PGConnect = applySecretUpdate(config.PGConnect, form.pgconnect, form.clear_pgconnect);
}

function applyInfluxConfigForm(body) {
  const form = parseRequestBody(body);

  if (typeof form.addr === 'string') {
    config.InfluxAddr = form.addr;
  }
  config.InfluxToken = applySecretUpdate(config.InfluxToken, form.token, form.clear_influx_token);
  if (typeof form.org === 'string') {
    config.InfluxOrg = form.org;
  }
  if (typeof form.bucket === 'string') {
    config.InfluxBucket = form.bucket;
  }
  config.InfluxEnable = form.enable === 'on';
  config.InfluxSkipTLS = form.skip === 'on';
}

function applyPrometheusConfigForm(body) {
  const form = parseRequestBody(body);

  config.PrometheusEnable = form.enable === 'on';
}

function applyRetentionConfigBody(body) {
  const params = parseRequestBody(body);
  const presenceRetention = Number(params.presenceRetention);
  const connectivityRetention = Number(params.connectivityRetention);

  if (!isPositiveIntegerValue(presenceRetention)) {
    return { error: 'invalid presenceRetention' };
  }
  if (!isPositiveIntegerValue(connectivityRetention)) {
    return { error: 'invalid connectivityRetention' };
  }

  config.TrimHist = presenceRetention;
  config.ConnectivityRetention = connectivityRetention;

  return { config: publicConfig() };
}

function historyFor(mac, datePrefix = '') {
  const hostEntry = fakeHosts.find((item) => item.Mac === mac) ?? fakeHosts[0];
  const rows = [];
  const baseDate = new Date('2026-08-23T10:15:00');

  for (let i = 0; i < 210; i += 1) {
    const sampleDate = new Date(baseDate.getTime() - i * config.Timeout * 1000);
    rows.push({
      ...hostEntry,
      ID: i + 1,
      Date: formatDate(sampleDate),
      Now: hostEntry.Now === 0 ? 0 : i % 13 === 0 ? 0 : 1,
    });
  }

  return datePrefix === '' ? rows : rows.filter((item) => item.Date.startsWith(datePrefix));
}

function findHostByID(id) {
  return fakeHosts.find((item) => item.ID === id);
}

function formatDate(date) {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  const hour = String(date.getHours()).padStart(2, '0');
  const minute = String(date.getMinutes()).padStart(2, '0');
  const second = String(date.getSeconds()).padStart(2, '0');

  return `${year}-${month}-${day} ${hour}:${minute}:${second}`;
}

function addActivity(hostEntry, eventType, options = {}) {
  activityEvents.push({
    ID: nextActivityId,
    HostID: hostEntry.ID,
    Mac: hostEntry.Mac,
    Name: hostEntry.Name,
    EventType: eventType,
    Date: formatDate(options.date ?? new Date()),
    IP: hostEntry.IP ?? '',
    Iface: hostEntry.Iface ?? '',
    DeviceType: hostEntry.DeviceType ?? '',
    OldValue: options.oldValue ?? '',
    NewValue: options.newValue ?? '',
  });
  nextActivityId += 1;
}

function addActivityMinutesAgo(hostEntry, eventType, minutesAgo, options = {}) {
  addActivity(hostEntry, eventType, {
    ...options,
    date: new Date(Date.now() - minutesAgo * 60000),
  });
}

function seedActivity() {
  const deletedHostSnapshot = {
    ID: 99,
    Name: 'old tablet',
    Iface: 'wifi0',
    IP: '192.168.1.70',
    Mac: 'AA:BB:CC:00:00:70',
    DeviceType: 'tablet',
  };

  addActivityMinutesAgo(fakeHosts[0], 'discovered', 1560);
  addActivityMinutesAgo(fakeHosts[1], 'discovered', 1515);
  addActivityMinutesAgo(fakeHosts[1], 'device-type-changed', 65, { oldValue: '', newValue: 'nas' });
  addActivityMinutesAgo(fakeHosts[0], 'known', 28);
  addActivityMinutesAgo(fakeHosts[4], 'discovered', 12);
  addActivityMinutesAgo(fakeHosts[3], 'offline', 8);
  addActivityMinutesAgo(fakeHosts[3], 'online', 2);
  addActivityMinutesAgo(deletedHostSnapshot, 'offline', 1440);

  for (let i = 0; i < 120; i += 1) {
    const hostEntry = i % 3 === 0 ? fakeHosts[3] : i % 3 === 1 ? fakeHosts[2] : fakeHosts[1];
    addActivityMinutesAgo(hostEntry, i % 2 === 0 ? 'online' : 'offline', 20 + i);
  }

  const changeTypes = ['discovered', 'known', 'unknown', 'device-type-changed'];
  for (let i = 0; i < 28; i += 1) {
    const hostEntry = fakeHosts[i % fakeHosts.length];
    const eventType = changeTypes[i % changeTypes.length];
    addActivityMinutesAgo(hostEntry, eventType, 90 + i * 3, {
      oldValue: eventType === 'device-type-changed' ? '' : undefined,
      newValue: eventType === 'device-type-changed' ? hostEntry.DeviceType : undefined,
    });
  }
}

function sortedActivityEvents() {
  return [...activityEvents].sort((left, right) => {
    const byDate = right.Date.localeCompare(left.Date);
    return byDate === 0 ? right.ID - left.ID : byDate;
  });
}

function parseActivityLimit(url) {
  const rawLimit = url.searchParams.get('limit') ?? '20';
  const limit = Number(rawLimit);
  return Number.isInteger(limit) && limit >= 1 && limit <= 100 ? limit : 0;
}

function parseActivityOffset(url) {
  const rawOffset = url.searchParams.get('offset') ?? '0';
  const offset = Number(rawOffset);
  return Number.isInteger(offset) && offset >= 0 ? offset : -1;
}

function parseActivityCategory(url) {
  const category = url.searchParams.get('category') ?? 'all';
  if (category === 'all') {
    return () => true;
  }
  if (category === 'connectivity') {
    return (event) => connectivityEvents.has(event.EventType);
  }
  if (category === 'changes') {
    return (event) => changeEvents.has(event.EventType);
  }

  return null;
}

function parseActivityEventTypeSet(url) {
  const values = url.searchParams.getAll('eventType');
  if (values.length === 0) {
    return { set: null };
  }

  for (const value of values) {
    if (!validActivityEvents.has(value)) {
      return { error: 'invalid eventType' };
    }
  }

  return { set: new Set(values) };
}

function activityMacSet(url) {
  const macs = url.searchParams.getAll('mac').filter(Boolean);
  return macs.length === 0 ? null : new Set(macs);
}

function activityFor(url, predicate = () => true) {
  const limit = parseActivityLimit(url);
  if (limit === 0) {
    return { error: 'invalid limit' };
  }

  const offset = parseActivityOffset(url);
  if (offset < 0) {
    return { error: 'invalid offset' };
  }

  const categoryPredicate = parseActivityCategory(url);
  if (categoryPredicate === null) {
    return { error: 'invalid category' };
  }

  const eventTypeFilter = parseActivityEventTypeSet(url);
  if (eventTypeFilter.error) {
    return { error: eventTypeFilter.error };
  }

  return sortedActivityEvents()
    .filter(categoryPredicate)
    .filter((event) => eventTypeFilter.set === null || eventTypeFilter.set.has(event.EventType))
    .filter(predicate)
    .slice(offset, offset + limit);
}

function activityStatsFor(url) {
  const macs = activityMacSet(url);
  const stats = {
    Total: 0,
    Online: 0,
    Offline: 0,
    Discovered: 0,
    Known: 0,
    Unknown: 0,
    DeviceTypeChanged: 0,
  };

  for (const event of activityEvents) {
    if (macs !== null && !macs.has(event.Mac)) {
      continue;
    }

    stats.Total += 1;
    if (event.EventType === 'online') stats.Online += 1;
    if (event.EventType === 'offline') stats.Offline += 1;
    if (event.EventType === 'discovered') stats.Discovered += 1;
    if (event.EventType === 'known') stats.Known += 1;
    if (event.EventType === 'unknown') stats.Unknown += 1;
    if (event.EventType === 'device-type-changed') stats.DeviceTypeChanged += 1;
  }

  return stats;
}

function activityDeviceOptions() {
  const seen = new Set();
  const options = [];

  for (const hostEntry of fakeHosts) {
    if (!hostEntry.Mac || seen.has(hostEntry.Mac)) {
      continue;
    }

    seen.add(hostEntry.Mac);
    options.push({
      HostID: hostEntry.ID,
      Mac: hostEntry.Mac,
      Name: hostEntry.Name,
      DeviceType: hostEntry.DeviceType,
      Exists: true,
    });
  }

  for (const event of sortedActivityEvents()) {
    if (!event.Mac || seen.has(event.Mac)) {
      continue;
    }

    seen.add(event.Mac);
    options.push({
      HostID: event.HostID,
      Mac: event.Mac,
      Name: event.Name,
      DeviceType: event.DeviceType,
      Exists: false,
    });
  }

  return options;
}

function routeReadOnly(req, res, url) {
  const pathname = decodeURIComponent(url.pathname);

  if (req.method === 'GET' && pathname === '/api/config') {
    sendJSON(res, publicConfig());
    return true;
  }

  if (req.method === 'GET' && pathname === '/api/version') {
    sendJSON(res, config.Version);
    return true;
  }

  if (req.method === 'GET' && pathname === '/api/all') {
    sendJSON(res, fakeHosts);
    return true;
  }

  if (req.method === 'GET' && pathname === '/api/activity/stats') {
    sendJSON(res, activityStatsFor(url));
    return true;
  }

  if (req.method === 'GET' && pathname === '/api/activity/devices') {
    sendJSON(res, activityDeviceOptions());
    return true;
  }

  if (req.method === 'GET' && pathname === '/api/activity') {
    const macs = activityMacSet(url);
    const events = activityFor(url, (event) => macs === null || macs.has(event.Mac));
    if (events.error) {
      sendJSON(res, { error: events.error }, 400);
      return true;
    }

    sendJSON(res, events);
    return true;
  }

  if (req.method === 'GET' && pathname === '/api/history') {
    sendJSON(res, fakeHosts.flatMap((item) => historyFor(item.Mac)));
    return true;
  }

  const hostActivityMatch = pathname.match(/^\/api\/host\/(\d+)\/activity$/);
  if (req.method === 'GET' && hostActivityMatch) {
    const id = Number(hostActivityMatch[1]);
    const hostEntry = findHostByID(id);
    if (!hostEntry) {
      sendJSON(res, { error: 'invalid host id' }, 400);
      return true;
    }

    const events = activityFor(url, (event) => event.HostID === id);
    if (events.error) {
      sendJSON(res, { error: events.error }, 400);
      return true;
    }

    sendJSON(res, events);
    return true;
  }

  const hostMatch = pathname.match(/^\/api\/host\/(\d+)$/);
  if (req.method === 'GET' && hostMatch) {
    const id = Number(hostMatch[1]);
    const hostEntry = findHostByID(id);
    if (!hostEntry) {
      sendJSON(res, { error: 'invalid host id' }, 400);
      return true;
    }

    sendJSON(res, hostEntry);
    return true;
  }

  const historyMatch = pathname.match(/^\/api\/history\/([^/]+)\/?$/);
  if (req.method === 'GET' && historyMatch) {
    sendJSON(res, historyFor(historyMatch[1]));
    return true;
  }

  const historyDateMatch = pathname.match(/^\/api\/history\/([^/]+)\/(.+)$/);
  if (req.method === 'GET' && historyDateMatch) {
    sendJSON(res, historyFor(historyDateMatch[1], historyDateMatch[2]));
    return true;
  }

  const publicAssetPath = req.method === 'GET' ? getLocalPublicAsset(pathname) : '';
  if (publicAssetPath) {
    res.writeHead(200, {
      'content-type': getAssetContentType(publicAssetPath),
      'cache-control': 'no-store',
    });
    createReadStream(publicAssetPath).pipe(res);
    return true;
  }

  return false;
}

async function routeSafeAction(req, res, url) {
  const pathname = decodeURIComponent(url.pathname);

  if (req.method === 'GET' && pathname === '/api/notify_test') {
    sendText(res, 'mock notification skipped');
    return true;
  }

  if (req.method === 'GET' && pathname === '/api/rescan') {
    sendText(res, 'mock rescan skipped');
    return true;
  }

  if (req.method === 'GET' && pathname.startsWith('/api/edit/')) {
    const editMatch = pathname.match(/^\/api\/edit\/(\d+)\/([^/]*)(?:\/(.*))?$/);

    if (editMatch) {
      const id = Number(editMatch[1]);
      const name = editMatch[2];
      const action = editMatch[3] ?? '';
      const hostEntry = findHostByID(id);

      if (hostEntry) {
        const oldKnown = hostEntry.Known;
        hostEntry.Name = name;

        if (action === 'toggle') {
          hostEntry.Known = 1 - hostEntry.Known;
          if (oldKnown !== hostEntry.Known) {
            addActivity(hostEntry, hostEntry.Known === 1 ? 'known' : 'unknown');
          }
        }
      }
    }

    sendJSON(res, 'OK');
    return true;
  }

  const deviceTypeMatch = pathname.match(/^\/api\/host\/(\d+)\/type$/);
  if (req.method === 'PATCH' && deviceTypeMatch) {
    const id = Number(deviceTypeMatch[1]);
    const hostEntry = findHostByID(id);
    if (!hostEntry) {
      sendJSON(res, { error: 'invalid host id' }, 400);
      return true;
    }

    const body = await readBody(req);
    const params = parseRequestBody(body);
    if (!isDeviceType(params.deviceType)) {
      sendJSON(res, { error: 'invalid deviceType' }, 400);
      return true;
    }

    const oldDeviceType = hostEntry.DeviceType;
    hostEntry.DeviceType = params.deviceType;
    if (oldDeviceType !== hostEntry.DeviceType) {
      addActivity(hostEntry, 'device-type-changed', {
        oldValue: oldDeviceType,
        newValue: hostEntry.DeviceType,
      });
    }
    sendJSON(res, hostEntry);
    return true;
  }

  if (req.method === 'GET' && pathname.startsWith('/api/host/del/')) {
    sendJSON(res, 'OK');
    return true;
  }

  if (req.method === 'GET' && pathname.startsWith('/api/host/add/')) {
    sendJSON(res, fakeHosts[0]);
    return true;
  }

  if (req.method === 'GET' && pathname.startsWith('/api/wol/')) {
    sendJSON(res, true);
    return true;
  }

  if (req.method === 'GET' && pathname.startsWith('/api/port/')) {
    sendJSON(res, false);
    return true;
  }

  if (req.method === 'POST' && pathname === '/api/config/color') {
    const body = await readBody(req);
    const params = parseRequestBody(body);
    const color = params.color ?? params.Color;

    if (!isColorMode(color)) {
      sendJSON(res, { error: 'invalid color' }, 400);
      return true;
    }

    config.Color = color;
    sendJSON(res, publicConfig());
    return true;
  }

  if (req.method === 'POST' && pathname === '/api/config/retention') {
    const body = await readBody(req);
    const result = applyRetentionConfigBody(body);
    if (result.error) {
      sendJSON(res, { error: result.error }, 400);
      return true;
    }

    sendJSON(res, result.config);
    return true;
  }

  if (req.method === 'POST' && pathname === '/api/config/') {
    const body = await readBody(req);
    applyBasicConfigForm(body);
    const referer = req.headers.referer || '/config';
    res.writeHead(303, { location: referer });
    res.end();
    return true;
  }

  if (req.method === 'POST' && pathname === '/api/config_settings/') {
    const body = await readBody(req);
    applySettingsConfigForm(body);
    const referer = req.headers.referer || '/config';
    res.writeHead(303, { location: referer });
    res.end();
    return true;
  }

  if (req.method === 'POST' && pathname === '/api/config_influx/') {
    const body = await readBody(req);
    applyInfluxConfigForm(body);
    const referer = req.headers.referer || '/config';
    res.writeHead(303, { location: referer });
    res.end();
    return true;
  }

  if (req.method === 'POST' && pathname === '/api/config_prometheus/') {
    const body = await readBody(req);
    applyPrometheusConfigForm(body);
    const referer = req.headers.referer || '/config';
    res.writeHead(303, { location: referer });
    res.end();
    return true;
  }

  if (req.method === 'POST' && pathname.startsWith('/api/config')) {
    await readBody(req);
    const referer = req.headers.referer || '/config';
    res.writeHead(303, { location: referer });
    res.end();
    return true;
  }

  return false;
}

const server = createServer(async (req, res) => {
  const url = new URL(req.url ?? '/', `http://${host}:${port}`);

  if (routeReadOnly(req, res, url)) {
    return;
  }

  if (await routeSafeAction(req, res, url)) {
    return;
  }

  sendJSON(res, { error: 'mock endpoint not found' }, 404);
});

seedActivity();

server.listen(port, host, () => {
  console.log(`LANventory mock API listening at http://${host}:${port}`);
});
