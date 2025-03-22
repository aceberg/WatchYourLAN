import { For, Show } from "solid-js"
import { appConfig } from "../../functions/exports"
import { apiPath } from "../../functions/api"

function Scan() {

  return (
    <div class="card border-primary">
      <div class="card-header">Scan settings</div>
      <div class="card-body table-responsive">
        <form action={apiPath + '/api/config_settings/'} method="post">
          <table class="table table-borderless"><tbody>
            <tr>
              <td>Interfaces</td>
              <td><input name="ifaces" type="text" class="form-control" value={appConfig().Ifaces}></input></td>
            </tr>
            <tr>
              <td>Timeout (seconds)</td>
              <td><input name="timeout" type="number" class="form-control" value={appConfig().Timeout}></input></td>
            </tr>
            <tr>
              <td>Args for arp-scan</td>
              <td><input name="arpargs" type="text" class="form-control" value={appConfig().ArpArgs}></input></td>
            </tr>
            <tr>
              <td>Arp Strings</td>
              <td>
                <For each={appConfig().ArpStrs}>{arpStr =>
                  <input name="arpstrs" type="text" class="form-control" value={arpStr}></input>
                }</For>
                <input name="arpstrs" type="text" class="form-control"></input>
              </td>
            </tr>
            <tr>
              <td>Log level</td>
              <td><select name="log" class="form-select">
              <For each={["debug","info","warn","error"]}>{level =>
                <Show
                  when={level == appConfig().LogLevel}
                  fallback={<option value={level}>{level}</option>}
                >
                <option value={level} selected>{level}</option>
                </Show>
              }</For>
              </select></td>
            </tr>
            <tr>
              <td>Trim History (hours)</td>
              <td><input name="trim" type="number" class="form-control" value={appConfig().TrimHist}></input></td>
            </tr>
            <tr>
              <td>Store History in DB</td>
              <td>
                <div class="form-check form-switch">
                {appConfig().HistInDB
                  ? <input class="form-check-input" type="checkbox" name="histdb" checked></input>
                  : <input class="form-check-input" type="checkbox" name="histdb"></input>
                }
                </div>
              </td>
            </tr>
            <tr>
              <td>Use DB</td>
              <td><select name="usedb" class="form-select">
                <Show
                  when={appConfig().UseDB == "sqlite"}
                  fallback={<>
                    <option value="sqlite">sqlite</option>
                    <option value="postgres" selected>postgres</option>
                  </>}
                >
                  <option value="sqlite" selected>sqlite</option>
                  <option value="postgres">postgres</option>
                </Show>
              </select></td>
            </tr>
            <tr>
              <td>PG Connect URL</td>
              <td>
                <textarea name="pgconnect" class="form-control" style="width: 100%;" rows="3" wrap="soft">{appConfig().PGConnect}</textarea>
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

export default Scan