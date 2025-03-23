import { apiDelHost, apiEditHost } from "../../functions/api";

function HostCard(_props: any) {

  let name:string = "";

  const handleInput = async (n: string) => {
    
    name = n;
    await apiEditHost(_props.host.ID, name, '');
  };

  const handleToggle = async () => {

    if (name == "") {
      name = _props.host.Name;
    }

    await apiEditHost(_props.host.ID, name, 'toggle');
  };

  const handleDel = async () => {
    
    await apiDelHost(_props.host.ID);
    window.location.href = '/';
  };

  return (
    <div class="card border-primary">
      <div class="card-header">Host</div>
      <div class="card-body table-responsive">
        <table class="table table-striped table-hover">
          <tbody>
          <tr>
            <td>ID</td>
            <td>{_props.host.ID}</td>
          </tr>
          <tr>
            <td>Name</td>
            <td>
            <input type="text" class="form-control" value={_props.host.Name}
              onInput={e => handleInput(e.target.value)}></input>
            </td>
          </tr>
          <tr>
            <td>DNS name</td>
            <td>{_props.host.DNS}</td>
          </tr>
          <tr>
            <td>Iface</td>
            <td>{_props.host.Iface}</td>
          </tr>
          <tr>
            <td>IP</td>
            <td>
              <a href={"http://" + _props.host.IP} target="_blank">{_props.host.IP}</a>
            </td>
          </tr>
          <tr>
            <td>MAC</td>
            <td>{_props.host.Mac}</td>
          </tr>
          <tr>
            <td>Hardware</td>
            <td>{_props.host.Hw}</td>
          </tr>
          <tr>
            <td>Date</td>
            <td>{_props.host.Date}</td>
          </tr>
          <tr>
            <td>Known</td>
            <td>
              <div class="form-check form-switch">
                <input class="form-check-input" type="checkbox" 
                    onClick={handleToggle}
                    checked={_props.host.Known == 1
                      ? true
                      : false
                    }
                    ></input>
              </div>
            </td>
          </tr>
          <tr>
            <td>Online</td>
            <td>{_props.host.Now == 1
                  ? <i class="bi bi-check-circle-fill" style="color:var(--bs-success);"></i>
                  : <i class="bi bi-circle-fill" style="color:var(--bs-gray-500);"></i>
            }</td>
          </tr>
          </tbody>
        </table>
        <button type="button" onClick={handleDel} class="btn btn-outline-danger">Delete host</button>
      </div>
    </div>
  )
}

export default HostCard