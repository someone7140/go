import { InputState } from "./input_state.js";

const state = new InputState("#submitForm");
const go = new Go();

WebAssembly.instantiateStreaming(
  fetch("../followCheckWasm.wasm"),
  go.importObject
).then((result) => {
  go.run(result.instance);
  state.updateLoading(false);
  const onSubmitClick = () => {
    console.log(
      checkInstagram(state.getInput().userId, state.getInput().password)
    );
  };
  state.settingOnButtonClick(onSubmitClick);
});
