.PHONY: submodule.ignore
submodule.ignore:
	@echo "--- Running task: $@ ---"
	@echo "Configuring all submodules to ignore dirty worktrees..."
	@git submodule foreach 'git config -f "$$toplevel/.git/config" submodule."$$path".ignore dirty'
	@echo "--- Task $@ finished. ---"
	@echo ""

.PHONY: submodule.clean
submodule.clean:
	@echo "--- Running task: $@ ---"
	@echo "Deinitializing all submodules..."
	@git submodule deinit -f --all
	@echo "--- Task $@ finished. ---"
	@echo ""

.PHONY: submodule.update-all
submodule.update-all:
	@echo "--- Running task: $@ ---"
	@echo "Updating all submodules..."
	@git submodule update --init --recursive
	@echo "--- Task $@ finished. ---"
	@echo ""