<template>
  <v-card class="elevation-12">
    <v-toolbar color="primary" dark flat>
      <v-toolbar-title>Login form {{ isAuth }}</v-toolbar-title>
      <v-spacer />
    </v-toolbar>
    <v-card-text>
      <v-form>
        <v-text-field
          label="Login"
          name="login"
          placeholder="Login by username or email"
          prepend-icon="person"
          type="text"
          v-model="username"
        />

        <v-text-field
          id="password"
          label="Password"
          placeholder="Enter your password"
          name="password"
          prepend-icon="lock"
          type="password"
          v-model="password"
        />
      </v-form>
    </v-card-text>
    <v-card-actions>
      <v-spacer />
      <v-btn color="primary" @click="submit">Login</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import { mapState, mapActions } from 'vuex';

export default {
  data: () => ({
    username: '',
    password: ''
  }),
  computed: {
    ...mapState('auth', {
      isAuth: state => state.authenticated
    })
  },
  methods: {
    submit() {
      this.doLogin({ username: this.username, password: this.password });
    },
    ...mapActions('auth', ['doLogin'])
  }
};
</script>
