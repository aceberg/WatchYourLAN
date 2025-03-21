import { apiDelHost, apiEditHost } from "../../functions/api";
import { currentHost } from "../../functions/exports";

function HostCard() {

  let name:string = "";

  const handleInput = async (n: string) => {
    
    name = n;
    await apiEditHost(currentHost().ID, name, '');
  };

  const handleToggle = async () => {

    if (name == "") {
      name = currentHost().Name;
    }

    await apiEditHost(currentHost().ID, name, 'toggle');
  };

  const handleDel = async () => {
    
    await apiDelHost(currentHost().ID);
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
            <td>{currentHost().ID}</td>
          </tr>
          <tr>
            <td>Name</td>
            <td>
            <input type="text" class="form-control" value={currentHost().Name}
              onInput={e => handleInput(e.target.value)}></input>
            </td>
          </tr>
          <tr>
            <td>DNS name</td>
            <td>{currentHost().DNS}</td>
          </tr>
          <tr>
            <td>Iface</td>
            <td>{currentHost().Iface}</td>
          </tr>
          <tr>
            <td>IP</td>
            <td>
              <a href={"http://" + currentHost().IP} target="_blank">{currentHost().IP}</a>
            </td>
          </tr>
          <tr>
            <td>MAC</td>
            <td>{currentHost().Mac}</td>
          </tr>
          <tr>
            <td>Hardware</td>
            <td>{currentHost().Hw}</td>
          </tr>
          <tr>
            <td>Date</td>
            <td>{currentHost().Date}</td>
          </tr>
          <tr>
            <td>Known</td>
            <td>
              <div class="form-check form-switch">
                <input class="form-check-input" type="checkbox" 
                    onClick={handleToggle}
                    checked={currentHost().Known == 1
                      ? true
                      : false
                    }
                    ></input>
              </div>
            </td>
          </tr>
          <tr>
            <td>Online</td>
            <td>{currentHost().Now == 1
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