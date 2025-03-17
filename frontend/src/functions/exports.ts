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

export interface Conf {
	Color:   string;
  Timeout: number;
};

export const [allHosts, setAllHosts] = createSignal<Host[]>([]);

export const [appConfig, setAppConfig] = createSignal<Conf>();

export const [editNames, setEditNames] = createSignal(false);

// export const [store, setStore] = createStore({ hosts: allHosts() });

// export const [updBody, setUpdBody] = createSignal(false);