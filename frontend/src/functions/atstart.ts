import { apiGetAllHosts, apiGetConfig } from "./api";
import { appConfig, type Conf, type Host, ifaces as savedIfaces, setAllHosts, setAppConfig, setBkpHosts, setIfaces } from "./exports";
import { filterAtStart, filterFunc } from "./filter";
import { sortAtStart } from "./sort";

export function runAtStart() {
  getHosts();
  filterFunc("ID", 0); // reset filter

  setInterval(() => {
    getHosts();
  }, 60000); // 60000 ms = 1 minute
}

export async function getHosts() {
  const [hosts, config] = await Promise.all([
    apiGetAllHosts(),
    apiGetConfig().catch(() => null),
  ]);

  if (config !== null) {
    setAppConfig(config);
  }

  if (hosts !== null) {
    setAllHosts(hosts);
    setBkpHosts(hosts);

    listIfaces(hosts, config);
    sortAtStart();
    filterAtStart();
  }
}

function listIfaces(hosts: Host[], config: Conf | null) {
  const ifaceOptions: string[] = [];
  const addIface = (iface: string) => {
    const trimmed = iface.trim();
    if (trimmed !== "" && !ifaceOptions.includes(trimmed)) {
      ifaceOptions.push(trimmed);
    }
  };

  for (let host of hosts) {
    addIface(host.Iface);
  }

  const configuredIfaces = config?.Ifaces || appConfig().Ifaces;
  for (let iface of configuredIfaces.split(/\s+/)) {
    addIface(iface);
  }

  if (ifaceOptions.length > 0) {
    setIfaces(ifaceOptions);
    return;
  }

  if (savedIfaces().length > 0) {
    return;
  }

  setIfaces([]);
}
