<script type="text/JavaScript">
function pow(a, b) {
	var ans = 1;
	while (b > 0){
		ans = ans * a;
		b = b - 1;
	}
	return ans;
}
var x = pow(2, 3);
document.write(x);
</script>
