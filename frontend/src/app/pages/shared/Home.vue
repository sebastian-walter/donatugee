<template>
    <div>
        <h1 class="display-1 mb-4">How it works</h1>
        <div class="text-center mb-4">
            <icon name="tasks" scale="4" class="mb-2"></icon>
            <h3 class="headline mb-4">Find your challenge</h3>
            <v-btn to="/challenges" class="primary">Find challenges</v-btn>
        </div>
        <v-divider></v-divider>
        <div class="text-center mt-4 mb-4">
            <icon name="sign-in" scale="4" class="mb-2"></icon>
            <h3 class="headline mb-4">Sign up and apply for a challenge</h3>
            <v-btn to="/refugee/register" class="primary">Sign up</v-btn>
        </div>
        <v-divider></v-divider>
        <div class="text-center mt-4 pb-4">
            <icon name="terminal" scale="4" class="mb-2"></icon>
            <h3 class="headline mb-4">Get your hardware and start hacking</h3>
        </div>
        <v-card class="mb-4">
            <v-card-title>
                <h3 class="headline">You are a company?</h3>
            </v-card-title>
            <v-card-text class="pt-0">
                <p>Register your company, arrange the donations and create a challenge.</p>
                <v-btn to="/company/register" class="primary">Register my company</v-btn>
            </v-card-text>
        </v-card>
    </div>
</template>

<script>
    export default {
    	name: 'Home',
    }
</script>