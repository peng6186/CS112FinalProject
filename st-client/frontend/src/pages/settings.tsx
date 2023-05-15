import * as React from 'react';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';
import Switch from '@mui/material/Switch';
import {FormControl, InputLabel, MenuItem, Select, SelectChangeEvent, TextField} from "@mui/material";
import {SetProxyMode} from "../../wailsjs/go/main/App";
import {useDispatch, useSelector} from "react-redux";
import {RootState} from "../store/store";
import {setLocalPort, setMode, setProxyAddress} from "../store/settingsSlice";

export function Settings() {
    // const [mode, setMode] = React.useState('direct');
    // const [localPort, setLocalPort] = React.useState('18888');
    // const [ProxyAddress, setProxyAddress] = React.useState(':8888');
    const isConnected = useSelector((state: RootState) => state.status.isConnected);
    const mode = useSelector((state:RootState) => state.settings.mode);
    const localPort = useSelector((state:RootState) => state.settings.localPort);
    const proxyAddress = useSelector((state:RootState) => state.settings.proxyAddress);

    const dispatch = useDispatch();

    const handleModeChange = (event: SelectChangeEvent) => {
        dispatch(setMode(event.target.value));
        SetProxyMode(event.target.value);
    };

    return (
        <List
            // sx={{ width: '100%'}}
            // subheader={<ListSubheader>Settings</ListSubheader>}
        >
            <ListItem>
                <ListItemText id="option-list-label-proxy-mode" primary="Mode" />
                <Select
                    value={mode}
                    onChange={handleModeChange}
                    size={"small"}
                >
                    <MenuItem value={"direct"}>Direct</MenuItem>
                    <MenuItem value={"proxy"}>Proxy</MenuItem>
                    {/*<MenuItem value={30}>Rules</MenuItem>*/}
                </Select>


            </ListItem>
            <ListItem>
                <ListItemText id="option-list-label-local-port" primary="Local Port" />
                <TextField
                    hiddenLabel
                    value={localPort}
                    onChange={(e) => dispatch(setLocalPort(e.target.value))}
                    disabled={isConnected}
                    variant="outlined"
                    size="small"
                />
            </ListItem>
            <ListItem>
                <ListItemText id="option-list-label-server-addr" primary="Server Address" />
                <TextField
                    hiddenLabel
                    value={proxyAddress}
                    onChange={(e) => dispatch(setProxyAddress(e.target.value))}
                    disabled={isConnected}
                    variant="outlined"
                    size="small"
                />
            </ListItem>
            {/*<ListItem>*/}
            {/*    <ListItemText id="option-list-label-proxy-on" primary="On" />*/}
            {/*    <Switch*/}
            {/*        edge="end"*/}
            {/*        onChange={handleSwitchToggle('on')}*/}
            {/*        checked={checked.indexOf('on') !== -1}*/}
            {/*        inputProps={{*/}
            {/*            'aria-labelledby': 'option-list-label-proxy-on',*/}
            {/*        }}*/}
            {/*    />*/}
            {/*</ListItem>*/}
        </List>
    );
}