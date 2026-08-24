import { createSignal, Show } from "solid-js"
import { apiPath } from "../../functions/api"
import { saveErrorMessage, submitConfigForm } from "../../functions/configForms";
import { appConfig } from "../../functions/exports"

function Influx() {
  const [status, setStatus] = createSignal("");
  const [error, setError] = createSignal("");

  const handleSubmit = async (event: SubmitEvent) => {
    event.preventDefault();

    const form = event.currentTarget as HTMLFormElement;
    setStatus("");
    setError("");

    try {
      await submitConfigForm(form);
      setStatus("Saved");
    } catch (saveError) {
      setError(saveErrorMessage(saveError));
    }
  };

  return (
    <div class="card wyl-panel config-panel">
          <div class="card-header">InfluxDB2 config</div>
          <div class="card-body table-responsive">
            <form action={apiPath + '/api/config_influx/'} method="post" onSubmit={handleSubmit}>
              <table class="table table-borderless"><tbody>
                <tr>
                  <td class="config-field-label">Enable</td>
                  <td class="config-field-value">
                    <div class="form-check form-switch">
                      {appConfig().InfluxEnable
                        ? <input class="form-check-input" type="checkbox" name="enable" checked></input>
                        : <input class="form-check-input" type="checkbox" name="enable"></input>
                      }
                    </div>
                  </td>
                </tr>
                <tr>
                  <td class="config-field-label">Address</td>
                  <td class="config-field-value"><input name="addr" type="text" class="form-control" value={appConfig().InfluxAddr}></input></td>
                </tr>
                <tr>
                  <td class="config-field-label">Token</td>
                  <td class="config-field-value">
                    <input
                      name="token"
                      type="text"
                      class="form-control"
                      placeholder={appConfig().InfluxTokenConfigured ? "Configured - leave blank to keep current value" : ""}
                    ></input>
                    <Show when={appConfig().InfluxTokenConfigured}>
                      <label class="form-check config-secret-clear">
                        <input name="clear_influx_token" class="form-check-input" type="checkbox"></input>
                        <span class="form-check-label">Clear stored InfluxDB token</span>
                      </label>
                    </Show>
                    <div class="config-field-helper">Stored InfluxDB tokens are write-only and are not displayed after saving.</div>
                  </td>
                </tr>
                <tr>
                  <td class="config-field-label">Org</td>
                  <td class="config-field-value"><input name="org" type="text" class="form-control" value={appConfig().InfluxOrg}></input></td>
                </tr>
                <tr>
                  <td class="config-field-label">Bucket</td>
                  <td class="config-field-value"><input name="bucket" type="text" class="form-control" value={appConfig().InfluxBucket}></input></td>
                </tr>
                <tr>
                  <td class="config-field-label">Skip TLS verify</td>
                  <td class="config-field-value">
                    <div class="form-check form-switch">
                      {appConfig().InfluxSkipTLS
                        ? <input class="form-check-input" type="checkbox" name="skip" checked></input>
                        : <input class="form-check-input" type="checkbox" name="skip"></input>
                      }
                    </div>
                  </td>
                </tr>
                <tr>
                  <td class="config-action-cell">
                    <button type="submit" class="btn btn-sm wyl-button">Save InfluxDB</button>
                    <span class={"config-save-status" + (error() ? " config-save-error" : "")} role="status">
                      {error() || status()}
                    </span>
                  </td>
                  <td class="config-action-cell"></td>
                </tr>
              </tbody></table>
            </form>
          </div>
        </div>
  )
}

export default Influx
