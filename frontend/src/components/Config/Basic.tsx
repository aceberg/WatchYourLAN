import { For, Show } from "solid-js";
import { apiPath, apiTestNotify } from "../../functions/api"
import { appConfig } from "../../functions/exports"

function Basic() {

  const themes = ["cerulean", "cosmo", "cyborg", "darkly", "emerald", "flatly", "grass", "grayscale", "journal", "litera", "lumen", "lux", "materia", "minty", "morph", "ocean", "pulse", "quartz", "sand", "sandstone", "simplex", "sketchy", "slate", "solar", "spacelab", "superhero", "united", "vapor", "wood", "yeti", "zephyr"];

  const handleTestNotify = () => {
    apiTestNotify();
  };

  return (
    <div class="card border-primary">
      <div class="card-header">Basic config</div>
      <div class="card-body table-responsive">
        <form action={apiPath + '/api/config/'} method="post">
          <table class="table table-borderless">
          <tbody>
            <tr>
              <td>Host</td>
              <td><input name="host" type="text" class="form-control" value={appConfig().Host}></input></td>
            </tr>
            <tr>
              <td>Port</td>
              <td><input name="port" type="text" class="form-control" value={appConfig().Port}></input></td>
            </tr>
            <tr>
              <td>Theme</td>
              <td>
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
              </td>
            </tr>
            <tr>
               <td>Color mode</td>
               <td>
                <select name="color" class="form-select">
                <Show
                  when={appConfig().Color == "dark"}
                  fallback={<>
                    <option value="dark">dark</option>
                    <option value="light" selected>light</option>
                  </>}
                >
                  <option value="dark" selected>dark</option>
                  <option value="light">light</option>
                </Show>
                </select>
               </td>
            </tr>
            <tr>
              <td>Local node-bootstrap URL</td>
              <td><input name="node" type="text" class="form-control" value={appConfig().NodePath}></input></td>
            </tr>
            <tr>
              <td>Shoutrrr URL</td>
              <td>
                <textarea name="shout" class="form-control" style="width: 100%;" rows="3" wrap="soft">{appConfig().ShoutURL}</textarea>
              </td>
            </tr>
            <tr>
              <td><button type="submit" class="btn btn-primary">Save</button></td>
              <td><button onClick={handleTestNotify} type="button" style="float: right;" class="btn btn-info">Test notification</button></td>
              <td></td>
            </tr>
          </tbody>
          </table>
        </form>
      </div>
    </div>
  )
}

export default Basic