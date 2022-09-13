<script>
	import { goto } from '$app/navigation';

	let email = '';
	let password = '';
	let errors = null;

	async function submit(event) {
		const response = await post(`auth/login`, { email, password });

		// TODO handle network errors
		errors = response.errors;

		if (response.user) {
			$session.user = response.user;
			goto('/');
		}
	}
</script>

<svelte:head>
	<title>Sign in > Yamori</title>
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
					<button class="btn btn-lg btn-primary pull-xs-right" type="submit"> Sign in </button>
				</fieldset>
			</form>
		</div>
	</div>
</div>
