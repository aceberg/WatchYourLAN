import { createSignal, For, Show } from "solid-js"
import { appConfig } from "../../functions/exports"
import { apiPath } from "../../functions/api"
import { saveErrorMessage, submitConfigForm } from "../../functions/configForms";

function Scan() {
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
      <div class="card-header">Scan settings</div>
      <div class="card-body table-responsive">
        <form action={apiPath + '/api/config_settings/'} method="post" onSubmit={handleSubmit}>
          <table class="table table-borderless"><tbody>
            <tr class="config-subsection-row">
              <td colSpan={2}>Network discovery</td>
            </tr>
            <tr>
              <td class="config-field-label">Interfaces</td>
              <td class="config-field-value"><input name="ifaces" type="text" class="form-control" value={appConfig().Ifaces}></input></td>
            </tr>
            <tr>
              <td class="config-field-label">Scan interval</td>
              <td class="config-field-value">
                <div class="config-value-with-unit">
                  <input name="timeout" type="number" class="form-control" value={appConfig().Timeout}></input>
                  <span class="config-field-unit">seconds</span>
                </div>
                <div class="config-field-helper">Time between network scans.</div>
              </td>
            </tr>
            <tr>
              <td class="config-field-label">Args for arp-scan</td>
              <td class="config-field-value"><input name="arpargs" type="text" class="form-control" value={appConfig().ArpArgs}></input></td>
            </tr>
            <tr>
              <td class="config-field-label config-field-label-top">ARP Strings</td>
              <td class="config-field-value">
                <For each={appConfig().ArpStrs}>{arpStr =>
                  <input name="arpstrs" type="text" class="form-control" value={arpStr}></input>
                }</For>
                <input name="arpstrs" type="text" class="form-control"></input>
              </td>
            </tr>
            <tr>
              <td class="config-field-label">Log level</td>
              <td class="config-field-value"><select name="log" class="form-select">
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
            <tr class="config-subsection-row">
              <td colSpan={2}>Database</td>
            </tr>
            <tr>
              <td class="config-field-label">Use DB</td>
              <td class="config-field-value"><select name="usedb" class="form-select">
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
              <td class="config-field-label config-field-label-top">PG Connect URL</td>
              <td class="config-field-value">
                <textarea
                  name="pgconnect"
                  class="form-control"
                  style="width: 100%;"
                  rows="3"
                  wrap="soft"
                  placeholder={appConfig().PGConnectConfigured ? "Configured - leave blank to keep current value" : ""}
                ></textarea>
                <Show when={appConfig().PGConnectConfigured}>
                  <label class="form-check config-secret-clear">
                    <input name="clear_pgconnect" class="form-check-input" type="checkbox"></input>
                    <span class="form-check-label">Clear stored PostgreSQL connection URL</span>
                  </label>
                </Show>
                <div class="config-field-helper">Stored database connection URLs are write-only and are not displayed after saving.</div>
              </td>
            </tr>
            <tr>
              <td></td>
              <td class="config-action-cell">
                <button type="submit" class="btn btn-sm wyl-button">Save scan settings</button>
                <span class={"config-save-status" + (error() ? " config-save-error" : "")} role="status">
                  {error() || status()}
                </span>
                <div class="config-field-helper config-save-helper">Saving these settings restarts network scanning.</div>
              </td>
            </tr>
            </tbody></table>
        </form>
      </div>
    </div>
  )
}

export default Scan
