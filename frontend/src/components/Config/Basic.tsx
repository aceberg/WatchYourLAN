import { createSignal, For, Show } from "solid-js";
import { apiPath, apiTestNotify } from "../../functions/api"
import { saveErrorMessage, submitConfigForm } from "../../functions/configForms";
import { appConfig } from "../../functions/exports"
import { applyColorMode, cacheColorMode, isColorMode } from "../../functions/theme";

function Basic() {
  const [status, setStatus] = createSignal("");
  const [error, setError] = createSignal("");

  const themes = ["cerulean", "cosmo", "cyborg", "darkly", "emerald", "flatly", "grass", "grayscale", "journal", "litera", "lumen", "lux", "materia", "minty", "morph", "ocean", "pulse", "quartz", "sand", "sandstone", "simplex", "sketchy", "slate", "solar", "spacelab", "superhero", "united", "vapor", "wood", "yeti", "zephyr"];

  const handleTestNotify = () => {
    apiTestNotify();
  };

  const handleSubmit = async (event: SubmitEvent) => {
    event.preventDefault();

    const form = event.currentTarget as HTMLFormElement;
    const color = new FormData(form).get("color");
    const previousColor = appConfig().Color;

    setStatus("");
    setError("");

    if (isColorMode(color)) {
      applyColorMode(color);
      cacheColorMode(color);
    }

    try {
      await submitConfigForm(form);
      setStatus("Saved");
    } catch (saveError) {
      if (isColorMode(previousColor)) {
        applyColorMode(previousColor);
        cacheColorMode(previousColor);
      }
      setError(saveErrorMessage(saveError));
    }
  };

  return (
    <div class="card wyl-panel config-panel">
      <div class="card-header">General</div>
      <div class="card-body table-responsive">
        <form action={apiPath + '/api/config/'} method="post" onSubmit={handleSubmit}>
          <table class="table table-borderless">
          <tbody>
            <tr class="config-subsection-row">
              <td colSpan={2}>Server</td>
            </tr>
            <tr>
              <td class="config-field-label">Host</td>
              <td class="config-field-value"><input name="host" type="text" class="form-control" value={appConfig().Host}></input></td>
            </tr>
            <tr>
              <td class="config-field-label">Port</td>
              <td class="config-field-value"><input name="port" type="text" class="form-control" value={appConfig().Port}></input></td>
            </tr>
            <tr class="config-subsection-row">
              <td colSpan={2}>Appearance</td>
            </tr>
            <tr>
              <td class="config-field-label">Base theme</td>
              <td class="config-field-value">
                <select name="theme" class="form-select">
                <For each={themes}>{theme =>
                  <Show
                    when={theme == appConfig().Theme}
                    fallback={<option value={theme}>{theme}</option>}
                  >
                    <option value={theme} selected>{theme}</option>
                  </Show>
                }</For>
                </select>
                <div class="config-field-helper">Base Bootstrap/Bootswatch styling underneath the LANventory interface.</div>
              </td>
            </tr>
            <tr>
               <td class="config-field-label">Color mode</td>
               <td class="config-field-value">
                <select name="color" class="form-select" value={appConfig().Color || "dark"}>
                  <option value="dark">dark</option>
                  <option value="light">light</option>
                </select>
               </td>
            </tr>
            <tr class="config-subsection-row">
              <td colSpan={2}>Notifications / compatibility</td>
            </tr>
            <tr>
              <td class="config-field-label config-field-label-top">Shoutrrr URL</td>
              <td class="config-field-value">
                <textarea
                  name="shout"
                  class="form-control"
                  style="width: 100%;"
                  rows="3"
                  wrap="soft"
                  placeholder={appConfig().ShoutURLConfigured ? "Configured - leave blank to keep current value" : ""}
                ></textarea>
                <Show when={appConfig().ShoutURLConfigured}>
                  <label class="form-check config-secret-clear">
                    <input name="clear_shout" class="form-check-input" type="checkbox"></input>
                    <span class="form-check-label">Clear stored Shoutrrr URL</span>
                  </label>
                </Show>
                <div class="config-field-helper">Stored notification URLs are write-only and are not displayed after saving.</div>
                <div class="config-inline-actions">
                  <button onClick={handleTestNotify} type="button" class="btn btn-sm wyl-button config-secondary-action">Test notification</button>
                </div>
              </td>
            </tr>
            <tr>
              <td class="config-field-label config-field-label-top">Local node-bootstrap URL</td>
              <td class="config-field-value">
                <input name="node" type="text" class="form-control" value={appConfig().NodePath}></input>
                <div class="config-field-helper">Legacy compatibility setting. UI theme assets are bundled locally by LANventory.</div>
              </td>
            </tr>
            <tr>
              <td></td>
              <td class="config-action-cell">
                <button type="submit" class="btn btn-sm wyl-button">Save general settings</button>
                <span class={"config-save-status" + (error() ? " config-save-error" : "")} role="status">
                  {error() || status()}
                </span>
              </td>
            </tr>
          </tbody>
          </table>
        </form>
      </div>
    </div>
  )
}

export default Basic
