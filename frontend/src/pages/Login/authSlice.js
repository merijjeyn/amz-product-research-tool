import { createSlice } from '@reduxjs/toolkit'

export const authSlice = createSlice({
    name: 'auth',
    initialState: {
        loggedIn: false
    },
    reducers: {
        loginAction: state => {
            state.loggedIn = true
        },
        logoutAction: state => {
            state.loggedIn = false
        }
    }
})

export const { loginAction, logoutAction } = authSlice.actions

export default authSlice.reducer
