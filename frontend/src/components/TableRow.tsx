


function TableRow(_props: any) {

  const link = "http://" + _props.host.IP;
  
  let now = <i class="bi bi-circle-fill" style="color:var(--bs-gray-500);"></i>;
  if (_props.host.Now == 1) {
    now = <i class="bi bi-check-circle-fill" style="color:var(--bs-success);"></i>;
  };

  let known = false;
  if (_props.host.Known == 1) {
    known = true;
  }

  return (
    <tr>
      <td class="opacity-50">{_props.index}.</td>
      <td>{_props.host.Name}</td>
      <td>{_props.host.Iface}</td>
      <td><a href={link} target="_blank">{_props.host.IP}</a></td>
      <td>{_props.host.Mac}</td>
      <td>{_props.host.Hw}</td>
      <td>{_props.host.Date}</td>
      <td>
        <div class="form-check form-switch">
          <input class="form-check-input" type="checkbox" checked={known}></input>
        </div>
      </td>
      <td>{now}</td>
      <td><i class="bi bi-three-dots-vertical my-btn"></i></td>
    </tr>
  )
}

export default TableRow
