<script>
  import { session } from '$app/stores'
  import { goto } from '$app/navigation'

</script>

<svelte:head>
  <title> Sign in > Yamori </title>
</svelte:head>
  
<div class="auth-page">
  <div class="container">
    <div class="column">
      <div class="header">
        <h1>Sign in</h1>
        <p>Need an account? <a href="/register">Sign up</a></p>
      </div>
      <form>
        <fieldset>
          <fieldset class="form-group">
            <input type="email" placeholder="Email" />
          </fieldset>
          <fieldset class="form-group">
            <input type="password" placeholder="Password" />
          </fieldset>
          <button class="btn btn-lg btn-primary pull-xs-right" type="submit">
            Sign in
          </button>
        </fieldset>
      </form>
    </div>
  </div>
</div>