import React from "react";
import { Button, client, TextField, useWizardContext } from "@clutch-sh/core";
import { useDataLayout } from "@clutch-sh/data-layout";
import { Wizard, WizardStep } from "@clutch-sh/wizard";
import { Typography } from "@mui/material";
const EchoInput = () => {
  const { onSubmit } = useWizardContext();
  const echoInput = useDataLayout("echoInput");
  return /* @__PURE__ */ React.createElement(WizardStep, null, /* @__PURE__ */ React.createElement(
    TextField,
    {
      InputLabelProps: { color: "secondary" },
      label: "Enter in text to echo back",
      name: "name",
      required: true,
      onChange: (e) => echoInput.assign({ [e.target.name]: e.target.value }),
      onReturn: onSubmit
    }
  ), /* @__PURE__ */ React.createElement(Button, { text: "Echo", destructive: true, onClick: onSubmit }));
};
const EchoOutput = () => {
  var _a, _b;
  const echoData = useDataLayout("echoOutput");
  return /* @__PURE__ */ React.createElement(WizardStep, { error: echoData.error, isLoading: echoData.isLoading }, /* @__PURE__ */ React.createElement(Typography, null, (_b = (_a = echoData.displayValue()) == null ? void 0 : _a.data) == null ? void 0 : _b.message));
};
const EchoWorkflow = ({ heading }) => {
  const dataLayout = {
    echoInput: {},
    echoOutput: {
      deps: ["echoInput"],
      hydrator: (echoInput) => {
        return client.post("/v1/echo/sayHello", {
          name: echoInput.name
        }).then((resp) => {
          return resp;
        });
      }
    }
  };
  return /* @__PURE__ */ React.createElement(Wizard, { dataLayout, heading }, /* @__PURE__ */ React.createElement(EchoInput, { name: "Lookup" }), /* @__PURE__ */ React.createElement(EchoOutput, { name: "Echo", heading: "Instance Details" }));
};
export default EchoWorkflow;
