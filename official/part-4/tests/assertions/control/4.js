<script type="text/JavaScript">
var count=0;
var oddeven = 0;
while (count < 10) {
	count = count + 1;
	if (oddeven == 0) {
		var a=count*2;
		oddeven = 1;
	} else {
		var b=count+1;
		oddeven = 0;
	}
	assert(a<14);
}
</script>
