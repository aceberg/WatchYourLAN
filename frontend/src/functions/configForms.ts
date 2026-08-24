import { refreshAppConfig } from "./theme";

export async function submitConfigForm(form: HTMLFormElement) {
  const body = new URLSearchParams();
  new FormData(form).forEach((value, key) => {
    if (typeof value === "string") {
      body.append(key, value);
    }
  });

  const response = await fetch(form.action, {
    method: form.method || "POST",
    headers: { "content-type": "application/x-www-form-urlencoded" },
    body,
  });

  if (!response.ok) {
    const detail = await response.text();
    throw new Error(detail || response.statusText || "Settings could not be saved.");
  }

  return await refreshAppConfig();
}

export function saveErrorMessage(error: unknown) {
  if (error instanceof Error && error.message.trim() !== "") {
    return error.message;
  }

  return "Settings could not be saved.";
}
