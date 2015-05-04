<script type="text/JavaScript">
function fib(a) {
	if(a == 0) {
		return 0;
	}
	if(a == 1) {
		return 1;
	}
	return fib(a - 1) + fib(a - 2);
}
document.write(fib(10));
</script>
