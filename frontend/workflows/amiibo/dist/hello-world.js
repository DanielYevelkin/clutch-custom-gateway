const React = require("react");
const { Button, ButtonGroup, Table, TableRow, TextField, client, useWizardContext } = require("@clutch-sh/core");
const { useDataLayout } = require("@clutch-sh/data-layout");
const { Wizard, WizardStep } = require("@clutch-sh/wizard");
const _ = require("lodash");

const AmiiboLookup = () => {
  const { onSubmit } = useWizardContext();
  const userInput = useDataLayout("userInput");

  const onChange = event => {
    userInput.assign({ name: event.target.value });
  };

  return React.createElement(
    React.Fragment,
    null,
    React.createElement(TextField, { onChange: onChange, onReturn: onSubmit }),
    React.createElement(
      ButtonGroup,
      null,
      React.createElement(Button, { text: "Search", onClick: onSubmit })
    )
  );
};

const AmiiboDetails = () => {
  const amiiboData = useDataLayout("amiiboData");
  let amiiboResults = amiiboData.displayValue();
  if (_.isEmpty(amiiboResults)) {
    amiiboResults = [];
  }

  return React.createElement(
    WizardStep,
    { error: amiiboData.error, isLoading: amiiboData.isLoading },
    React.createElement(
      Table,
      { columns: ["Name", "Image", "Series", "Type"] },
      amiiboResults.map(amiibo =>
        React.createElement(
          TableRow,
          { key: amiibo.imageUrl },
          amiibo.name,
          React.createElement("img", { alt: amiibo.name, src: amiibo.imageUrl, height: "75px" }),
          amiibo.amiiboSeries,
          amiibo.type
        )
      )
    )
  );
};

const HelloWorld = ({ heading }) => {
  const dataLayout = {
    userInput: {},
    amiiboData: {
      deps: ["userInput"],
      hydrator: userInput => {
        return client
          .post("/v1/amiibo/getAmiibo", {
            name: userInput.name,
          })
          .then(response => {
            return response?.data?.amiibo || [];
          });
      },
    },
  };

  return React.createElement(
    Wizard,
    { dataLayout: dataLayout, heading: heading },
    React.createElement(AmiiboLookup, { name: "Lookup" }),
    React.createElement(AmiiboDetails, { name: "Details" })
  );
};

module.exports = HelloWorld;
