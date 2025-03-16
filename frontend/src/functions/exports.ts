import { createSignal } from "solid-js";

export interface Host {
	ID:    number;
	Name:  string;
	DNS:   string;
	Iface: string;
	IP:    string;
	Mac:   string;
	Hw:    string;
	Date:  string;
	Known: number;
	Now:   number;
};

export const emptyHost: Host = {
	ID: 0,
	Name: "",
	DNS: "",
	Iface: "",
	IP:    "",
	Mac:   "",
	Hw:    "",
	Date:  "",
	Known: 0,
	Now:   0,
};

export const [allHosts, setAllHosts] = createSignal<Host[]>([]);