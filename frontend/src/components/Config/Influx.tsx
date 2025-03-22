import { apiPath } from "../../functions/api"
import { appConfig } from "../../functions/exports"

function Influx() {

  return (
    <div class="card border-primary">
          <div class="card-header">InfluxDB2 config</div>
          <div class="card-body table-responsive">
            <form action={apiPath + '/api/config_influx/'} method="post">
              <table class="table table-borderless"><tbody>
                <tr>
                  <td>Enable</td>
                  <td>
                    <div class="form-check form-switch">
                      {appConfig().InfluxEnable
                        ? <input class="form-check-input" type="checkbox" name="enable" checked></input>
                        : <input class="form-check-input" type="checkbox" name="enable"></input>
                      }
                    </div>
                  </td>
                </tr>
                <tr>
                  <td>Address</td>
                  <td><input name="addr" type="text" class="form-control" value={appConfig().InfluxAddr}></input></td>
                </tr>
                <tr>
                  <td>Token</td>
                  <td><input name="token" type="text" class="form-control" value={appConfig().InfluxToken}></input></td>
                </tr>
                <tr>
                  <td>Org</td>
                  <td><input name="org" type="text" class="form-control" value={appConfig().InfluxOrg}></input></td>
                </tr>
                <tr>
                  <td>Bucket</td>
                  <td><input name="bucket" type="text" class="form-control" value={appConfig().InfluxBucket}></input></td>
                </tr>
                <tr>
                  <td>Skip TLS verify</td>
                  <td>
                    <div class="form-check form-switch">
                      {appConfig().InfluxSkipTLS
                        ? <input class="form-check-input" type="checkbox" name="skip" checked></input>
                        : <input class="form-check-input" type="checkbox" name="skip"></input>
                      }
                    </div>
                  </td>
                </tr>
                <tr>
                  <td><button type="submit" class="btn btn-primary">Save</button></td>
                  <td></td>
                </tr>
              </tbody></table>
            </form>
          </div>
        </div>
  )
}

export default Influx