import './App.css';

import {
    ThemeProvider,
} from "@mui/material";
import CssBaseline from "@mui/material/CssBaseline";
import Box from "@mui/material/Box";
import * as React from "react";
import {BarChart, Dashboard, Settings} from "@mui/icons-material";

import {theme} from "./theme";
import {Nav} from "./components/nav/nav";
import {Outlet} from "react-router-dom";


const data = [
    { icon: <Dashboard />, label: 'Dashboard', nav: '/' },
    // { icon: <BarChart />, label: 'Connections', nav: '/connections' },
    { icon: <Settings />, label: 'Settings', nav: '/settings' },
];




function App() {
    const [value, setValue] = React.useState(0);

    const handleChange = (event: React.SyntheticEvent, newValue: number) => {
        setValue(newValue);
    };


    return (
        <Box id="App">
            <ThemeProvider theme={theme}>
            <Box sx={{ display: 'flex' }}>
                <CssBaseline />
                <Nav data={data}/>
                <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
                    <Outlet />
                </Box>
            </Box>
            </ThemeProvider>
        </Box>
    )
}

export default App
