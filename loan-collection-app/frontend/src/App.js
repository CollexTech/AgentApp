import React from "react";
import { Routes, Route } from "react-router-dom";
import LoginForm from "./components/LoginForm";
import CasesList from "./components/CasesList";
import CaseDetails from "./components/CaseDetails";
import { createTheme, ThemeProvider } from "@mui/material/styles";

const theme = createTheme({
  palette: {
    primary: {
      main: "#1976d2",
    },
  },
});

function App() {
  return (
    <ThemeProvider theme={theme}>
      <Routes>
        <Route path="/" element={<LoginForm />} />
        <Route path="/cases" element={<CasesList />} />
        <Route path="/cases/:caseId" element={<CaseDetails />} />
      </Routes>
    </ThemeProvider>
  );
}

export default App;
