<template>
  <v-row align="center" justify="center">
    <v-col cols="12" sm="8" md="4">
      <LoginForm />
    </v-col>
  </v-row>
</template>

<script>
import LoginForm from '@/components/auth/LoginForm.vue';
export default {
  name: 'Home',
  components: { LoginForm }
};
</script>
