<script>
	import { goto } from '$app/navigation';
  import loginSession from '$lib/stores';


  let user = loginSession 
  
  let errors = null;


  const unsbscribe = user.subscribe((value) => {
    
    console.log('value', value);
    if (value) {
      errors = value.errors;
    }
  });

  const onSubmit = async () => {
    try {
      const res = await fetch('/api/users/login', {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(user)
      });

      if (res.ok) {
        const { user } = await res.json();
        session.set(user);
        goto('/');
      } else {
        const { errors } = await res.json();
        throw errors;
      }
    } catch (err) {
      errors = err;
    }
  };

  
</script>
 

<h1>Sign in</h1>
<p>
  <a href="/register">Need an account?</a>
</p>
<form on:submit>
  <fieldset>
    <fieldset class="form-group">
      <input
        class="form-control form-control-lg"
        type="email"
        placeholder="Email"
        bind:value={$user.email}
      />
    </fieldset>
    <fieldset class="form-group">
      <input
        class="form-control form-control-lg"
        type="password"
        placeholder="Password"
        bind:value={$user.password}
      />
    </fieldset>
    <button class="btn btn-lg btn-primary pull-xs-right" type="submit" disabled={inProgress}>
      Sign in
    </button>
  </fieldset>
</form>
