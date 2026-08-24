import { createEffect, createSignal, For, onCleanup, Show } from "solid-js";
import { deviceTypes, getDeviceTypeOption, type DeviceTypeOption, type DeviceTypeValue } from "../functions/deviceTypes";

type DeviceTypePickerProps = {
  value: string | null | undefined;
  onChange: (value: DeviceTypeValue) => Promise<void> | void;
  mode?: "icon" | "full";
  class?: string;
  ariaLabel?: string;
  disabled?: boolean;
};

function DeviceTypePicker(props: DeviceTypePickerProps) {
  const [open, setOpen] = createSignal(false);
  const [saving, setSaving] = createSignal(false);
  const [saveError, setSaveError] = createSignal(false);
  let rootRef: HTMLDivElement | undefined;

  const mode = () => props.mode ?? "icon";
  const disabled = () => props.disabled === true;
  const selected = () => getDeviceTypeOption(props.value);
  const rootClass = () => [
    "device-type-picker",
    "device-type-picker-" + mode(),
    props.class ?? "",
  ].filter(Boolean).join(" ");
  const triggerClass = () => [
    "device-type-trigger",
    "device-type-trigger-" + mode(),
    "device-type-tone-" + selected().tone,
    saveError() ? "has-error" : "",
  ].filter(Boolean).join(" ");
  const triggerLabel = () => selected().label;
  const triggerAriaLabel = () => props.ariaLabel ?? "Device type: " + selected().label;

  createEffect(() => {
    if (!open()) {
      return;
    }

    const handlePointerDown = (event: PointerEvent) => {
      if (rootRef && !rootRef.contains(event.target as Node)) {
        setOpen(false);
      }
    };

    document.addEventListener("pointerdown", handlePointerDown);
    onCleanup(() => document.removeEventListener("pointerdown", handlePointerDown));
  });

  const handleSelect = async (option: DeviceTypeOption) => {
    if (saving() || disabled()) {
      return;
    }

    if (option.value === selected().value) {
      setOpen(false);
      return;
    }

    setSaving(true);
    setSaveError(false);
    try {
      await props.onChange(option.value);
      setOpen(false);
    } catch {
      setSaveError(true);
    } finally {
      setSaving(false);
    }
  };

  const handleKeyDown = (event: KeyboardEvent) => {
    if (event.key === "Escape") {
      event.preventDefault();
      setOpen(false);
    }
  };

  const handleTriggerKeyDown = (event: KeyboardEvent) => {
    if (event.key === "Enter" || event.key === " ") {
      event.preventDefault();
      if (disabled()) {
        return;
      }
      setOpen((value) => !value);
    }
  };

  return (
    <div
      ref={rootRef}
      class={rootClass()}
      onKeyDown={handleKeyDown}
      onClick={(event) => event.stopPropagation()}
    >
      <button
        type="button"
        class={triggerClass()}
        title={saveError() ? "Device type could not be saved" : triggerLabel()}
        aria-label={saveError() ? "Device type could not be saved" : triggerAriaLabel()}
        aria-haspopup="listbox"
        aria-expanded={open()}
        aria-busy={saving()}
        disabled={disabled()}
        onClick={() => setOpen((value) => !value)}
        onKeyDown={handleTriggerKeyDown}
      >
        <i class={"bi " + selected().icon} aria-hidden="true"></i>
        <Show when={mode() === "full"}>
          <span class="device-type-trigger-label">{selected().label}</span>
        </Show>
      </button>

      <Show when={open()}>
        <div class="device-type-menu" role="listbox" aria-label="Device type">
          <For each={deviceTypes}>{(option) =>
            <button
              type="button"
              class={option.value === selected().value ? "device-type-option is-selected" : "device-type-option"}
              role="option"
              aria-selected={option.value === selected().value}
              disabled={saving() || disabled()}
              onClick={() => handleSelect(option)}
            >
              <i class={"bi " + option.icon + " device-type-option-icon device-type-tone-" + option.tone} aria-hidden="true"></i>
              <span class="device-type-option-label">{option.label}</span>
              <Show when={option.value === selected().value}>
                <i class="bi bi-check-lg device-type-option-check" aria-hidden="true"></i>
              </Show>
            </button>
          }</For>
        </div>
      </Show>
    </div>
  );
}

export default DeviceTypePicker;
