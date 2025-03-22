import { apiPath } from "../../functions/api"
import { appConfig } from "../../functions/exports"

function Prometheus() {

  return (
    <div class="card border-primary">
      <div class="card-header">Prometheus config</div>
      <div class="card-body table-responsive">
        <form action={apiPath + '/api/config_prometheus/'} method="post">
          <table class="table table-borderless"><tbody>
            <tr>
              <td>Enable</td>
              <td>
                <div class="form-check form-switch">
                  {appConfig().PrometheusEnable
                    ? <input class="form-check-input" type="checkbox" name="enable" checked></input>
                    : <input class="form-check-input" type="checkbox" name="enable"></input>
                  }
                </div>
              </td>
            </tr>
            <tr>
              <td><button type="submit" class="btn btn-primary">Save</button></td>
              <td>
                <a href="/metrics" target="_self">/metrics</a>
              </td>
            </tr>
          </tbody></table>
        </form>
      </div>
    </div>
  )
}

export default Prometheus