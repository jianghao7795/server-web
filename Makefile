status:
	git status
diff:
	git diff
diff-cache:
	git diff --cached
commit:
	git add .
	git commit -m "${msg}"
push:
	git push origin ${local}:${remote}

dev:
	pnpm run dev