import { createSignal, Show } from "solid-js";
import { editNames } from "../functions/exports";
import { apiEditHost } from "../functions/api";

function TableRow(_props: any) {

  const [name, setName] = createSignal(_props.host.Name);

  const link = "http://" + _props.host.IP;
  const edit = "/host/" + _props.host.ID;
  
  let now = <i class="bi bi-circle-fill" style="color:var(--bs-gray-500);"></i>;
  if (_props.host.Now == 1) {
    now = <i class="bi bi-check-circle-fill" style="color:var(--bs-success);"></i>;
  };

  let known = false;
  if (_props.host.Known == 1) {
    known = true;
  }

  const handleInput = async (n: string) => {
    // console.log('Edit:', n);
    setName(n);
    await apiEditHost(_props.host.ID, name(), '');
  };
  const handleToggle = async () => {
    await apiEditHost(_props.host.ID, name(), 'toggle');
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
      <td><a href={link} target="_blank">{_props.host.IP}</a></td>
      <td>{_props.host.Mac}</td>
      <td>{_props.host.Hw}</td>
      <td>{_props.host.Date}</td>
      <td>
        <div class="form-check form-switch">
          <input class="form-check-input" type="checkbox" checked={known}
            onClick={handleToggle}></input>
        </div>
      </td>
      <td>{now}</td>
      <td>
        <a href={edit}><i class="bi bi-pencil-square my-btn"></i></a>
      </td>
    </tr>
  )
}

export default TableRow
