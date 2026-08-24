import { createSignal, Show } from "solid-js";
import { editNames, hasMultipleIfaces, selectedIDs, setSelectedIDs } from "../../functions/exports";
import { apiEditHost, apiSetDeviceType } from "../../functions/api";
import { getHosts } from "../../functions/atstart";
import { formatLastSeen } from "../../functions/dateFormat";
import { updateHostInView } from "../../functions/hostView";
import type { DeviceTypeValue } from "../../functions/deviceTypes";
import DeviceTypePicker from "../DeviceTypePicker";

import { debounce } from "@solid-primitives/scheduled"; 

function TableRow(_props: any) {

  const [name, setName] = createSignal(_props.host.Name);

  const isOnline = () => _props.host.Now === 1;
  const known = () => _props.host.Known === 1;
  const lastSeen = () => formatLastSeen(_props.host.Date);
  const rowClass = () => [
    _props.host.Known === 0 ? "device-row-unknown" : "",
    !isOnline() ? "device-row-offline" : "",
  ].filter(Boolean).join(" ");
  const nameClass = () => isOnline() ? "" : "device-offline-name";
  const knownTitle = () => known()
    ? "Known device - click to mark unknown"
    : "Unknown device - click to mark known";
  const statusText = () => isOnline() ? "Online" : "Offline";

  const debouncedApi = debounce(async (val: string) => {
    await apiEditHost(_props.host.ID, val, "");
  }, 300);

  const handleInput = async (n: string) => {
    setName(n);
    debouncedApi(n);
  };
  const handleToggle = async () => {
    await apiEditHost(_props.host.ID, name(), "toggle");
    await getHosts();
  };

  const handleDeviceTypeChange = async (deviceType: DeviceTypeValue) => {
    const updatedHost = await apiSetDeviceType(_props.host.ID, deviceType);
    updateHostInView(updatedHost);
  };

  const handleCheck = (checked: boolean) => {
    const id = _props.host.ID;
    setSelectedIDs(prev => {
      if (checked) {
        return prev.includes(id) ? prev : [...prev, id];
      } else {
        return prev.filter(item => item !== id);
      }
    });
  };

  return (
    <tr class={rowClass()}>
      <td class="device-table-index opacity-50">{_props.index}.</td>
      <td class="device-table-known">
        <button
          type="button"
          class={known() ? "device-known-toggle device-known-toggle-known" : "device-known-toggle device-known-toggle-unknown"}
          title={knownTitle()}
          aria-label={knownTitle()}
          aria-pressed={known()}
          onClick={handleToggle}
        >
          <i class={known() ? "bi bi-bookmark-check-fill" : "bi bi-question-circle-fill"} aria-hidden="true"></i>
        </button>
      </td>
      <td class="device-table-actions">
        <Show
          when={editNames()}
          fallback={
          <a href={"/host/" + _props.host.ID + "?edit=1"} class="device-action-link" title="Edit host">
            <i class="bi bi-pencil-fill my-btn p-2" aria-hidden="true"></i>
          </a>}
        >
          <input
            type="checkbox"
            class="form-check-input device-action-checkbox"
            checked={selectedIDs().includes(_props.host.ID)}
            onChange={e => handleCheck((e.target as HTMLInputElement).checked)}
          />
        </Show>
      </td>
      <td class="device-table-name">
        <Show
          when={editNames()}
          fallback={<a href={"/host/" + _props.host.ID} class={"device-name-link " + nameClass()}>{name()}</a>}
        >
          <input type="text" class="form-control" value={name()}
            onInput={e => handleInput(e.target.value)}></input>
        </Show>
      </td>
      <td class="device-table-type">
        <DeviceTypePicker
          value={_props.host.DeviceType}
          mode="icon"
          onChange={handleDeviceTypeChange}
        ></DeviceTypePicker>
      </td>
      <td class="device-table-ip">
        <span class="device-ip-with-status">
          <span
            class={isOnline() ? "device-status-icon device-status-icon-online" : "device-status-icon device-status-icon-offline"}
            title={statusText()}
            aria-label={statusText()}
            role="img"
          >
            <i class={isOnline() ? "bi bi-check-circle-fill" : "bi bi-x-circle-fill"} aria-hidden="true"></i>
          </span>
          <Show when={isOnline()} fallback={<span class="device-ip-offline">{_props.host.IP}</span>}>
            <a href={"http://" + _props.host.IP} target="_blank" rel="noreferrer">{_props.host.IP}</a>
          </Show>
        </span>
      </td>
      <Show when={hasMultipleIfaces()}>
        <td class="device-table-iface"><span class="device-cell-muted">{_props.host.Iface}</span></td>
      </Show>
      <td class="device-table-mac"><span class="device-cell-muted">{_props.host.Mac}</span></td>
      <td class="device-table-hardware" title={_props.host.Hw}>
        <span class="device-hardware-text">{_props.host.Hw}</span>
      </td>
      <td class="device-table-last-seen" title={_props.host.Date}>
        <span class="device-cell-muted">{lastSeen()}</span>
      </td>
    </tr>
  )
}

export default TableRow
