export const InputState = class {
  constructor(selector) {
    this.parent = document.querySelector(selector);
    this.inputState = {
      loadingFlag: true,
      userId: "",
      password: "",
    };

    const judgeSubmitEnable = () => {
      const state = this.inputState;
      const submitButton = document.querySelector("#submitButton");
      if (state.userId && state.password && !state.loadingFlag) {
        submitButton.disabled = null;
      } else {
        submitButton.disabled = "disabled";
      }
    };

    const changeEvent = (e) => {
      const target = e.target;
      const bindAttr = target.getAttribute("data-bind");
      this.inputState[bindAttr] = target.value;
      judgeSubmitEnable();
    };

    const allBind = this.parent.querySelectorAll("[data-bind]");
    allBind.forEach((item) => {
      item.addEventListener("input", changeEvent);
    });

    judgeSubmitEnable();
  }

  getInput = () => {
    return this.inputState;
  };

  updateLoading = (flag) => {
    this.inputState.loadingFlag = flag;
  };

  settingOnButtonClick = (onSubmitClick) => {
    const submitButton = document.querySelector("#submitButton");
    submitButton.onclick = onSubmitClick;
  };
};
