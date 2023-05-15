import {createSlice, PayloadAction} from "@reduxjs/toolkit";

export interface SettingsState {
    mode: string;
    localPort: string;
    proxyAddress: string;
}

const initialState: SettingsState = {
    mode: "direct",
    localPort: "18888",
    proxyAddress: "localhost:8888",
}

export const settingsSlice = createSlice({
    name: 'settings',
    initialState,
    reducers: {
        setMode: (state, action:PayloadAction<string>) => {
            state.mode = action.payload;
        },
        setLocalPort: (state, action:PayloadAction<string>) => {
            state.localPort = action.payload;
        },
        setProxyAddress: (state, action:PayloadAction<string>) => {
            state.proxyAddress = action.payload;
        }
    }
})

export const {setMode, setLocalPort, setProxyAddress} = settingsSlice.actions;

export default settingsSlice.reducer;