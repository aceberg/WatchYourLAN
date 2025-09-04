import { createSignal, Show } from "solid-js";
import { editNames } from "../../functions/exports";
import { apiEditHost } from "../../functions/api";

import { debounce } from "@solid-primitives/scheduled"; 

function TableRow(_props: any) {

  const [name, setName] = createSignal(_props.host.Name);
  
  let now = <i class="bi bi-circle-fill" style="color:var(--bs-gray-500);"></i>;
  if (_props.host.Now == 1) {
    now = <i class="bi bi-check-circle-fill" style="color:var(--bs-success);"></i>;
  };

  let known:boolean;
  _props.host.Known === 1 ? known = true : known = false;

  const debouncedApi = debounce((val: string) => {
    apiEditHost(_props.host.ID, val, "");
  }, 300);

  const handleInput = async (n: string) => {
    setName(n);
    debouncedApi(n);
  };
  const handleToggle = async () => {
    await apiEditHost(_props.host.ID, name(), "toggle");
  };

  return (
    <tr>
      <td class="opacity-50">{_props.index}.</td>
      <td>
        <Show
          when={editNames()}
          fallback={name()}
        >
          <input type="text" class="form-control" value={name()}
            onInput={e => handleInput(e.target.value)}></input>
        </Show>
      </td>
      <td>{_props.host.Iface}</td>
      <td><a href={"http://" + _props.host.IP} target="_blank">{_props.host.IP}</a></td>
      <td>{_props.host.Mac}</td>
      <td title={_props.host.Hw}>{_props.host.Hw.slice(0,12)+".."}</td>
      <td>{_props.host.Date}</td>
      <td>
        <div class="form-check form-switch">
          <input class="form-check-input" type="checkbox" checked={known}
            onClick={handleToggle}></input>
        </div>
      </td>
      <td>{now}</td>
      <td>
        <a href={"/host/" + _props.host.ID}>
          <i class="bi bi-three-dots-vertical my-btn p-2" title="More"></i>
        </a>
      </td>
    </tr>
  )
}

export default TableRow
