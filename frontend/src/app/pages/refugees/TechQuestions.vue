<template>
    <div>
        <v-card>
            <v-card-title primary-title>
                <div>
                    <h3 class="headline mb-0">{{ questions.question }}</h3>
                </div>
            </v-card-title>
        </v-card>
    </div>
</template>
<script>
    export default {
		components: {VCard},
		name: 'TechQuestions',

    }
</script>