import {Grid, Paper, styled, useTheme} from "@mui/material";
import LinkIcon from '@mui/icons-material/Link';
import LinkOffIcon from '@mui/icons-material/LinkOff';
import {useSelector} from "react-redux";
import {RootState} from "../store/store";

const DashboardHeader = styled('h2')(({theme}) => ({
    marginTop: '12px',
    marginLeft: '18px',
    textAlign: 'left',
}));



export function Dashboard() {
    const theme = useTheme();
    const isConnected = useSelector((state: RootState) => state.status.isConnected);
    return (
        // <Paper sx={{paddingY: '36px'}}>
        //     <h1>Dashboard</h1>
        //     <p>This is dashboard page.</p>
        // </Paper>
        <Grid container spacing={2}>
            <Grid item xs={4}>
                <Paper sx={{p: 2, display: 'flex', flexDirection: 'column', height: 240, textAlign: 'center', backgroundColor: isConnected ? theme.palette.success.dark: theme.palette.background.paper}}>
                    <DashboardHeader>Status</DashboardHeader>
                    {isConnected ? <LinkIcon sx={{fontSize: '48px', textAlign: 'center', width: '100%', height: '50%'}}/> : <LinkOffIcon sx={{fontSize: '48px', textAlign: 'center', width: '100%', height: '50%'}}/>}
                </Paper>
            </Grid>
            <Grid item xs={4}>
                <Paper sx={{p: 2, display: 'flex', flexDirection: 'column', height: 240}}>
                    <DashboardHeader>Speed</DashboardHeader>
                </Paper>
            </Grid>
            <Grid item xs={4}>
                <Paper sx={{p: 2, display: 'flex', flexDirection: 'column', height: 240}}>
                    <DashboardHeader>Traffic</DashboardHeader>
                </Paper>
            </Grid>
            <Grid item xs={12}>
                <Paper sx={{p: 2, display: 'flex', flexDirection: 'column', height: 295}}>

                </Paper>
            </Grid>
        </Grid>
    )
}