.PHONY: default submodule.ignore

default:
	@echo "--- Please specify a target. Available targets: ---"
	@echo "  submodule.ignore"
	@echo "  submodule.deinit-all"
	@echo "  submodule.update-all"
	@echo ""
	@echo "  Pattern targets (use: make <target>-<submodule_path>):"
	@echo "    submodule.deinit-%"
	@echo "    submodule.update-%"
	@echo "--------------------------------------------------"
	
submodule.ignore:
	@echo "--- Running task: $@ ---"
	@echo "Configuring all submodules to ignore dirty worktrees..."
	@git submodule foreach 'git config -f "$$toplevel/.git/config" submodule."$$path".ignore dirty'
	@echo "--- Task $@ finished. ---"
	@echo ""

.PHONY: submodule.clean submodule.deinit-all submodule.deinit-%

submodule.deinit-%:
	@echo "--- Running task: $@ ---"
	@echo "Deinitializing submodule: $*..."
	@git submodule deinit -f $* 
	@echo "--- Task $@ finished. ---"
	@echo ""

submodule.deinit-all:
	@echo "--- Running task: $@ ---"
	@echo "Deinitializing all submodules..."
	@git submodule deinit -f --all
	@echo "--- Task $@ finished. ---"
	@echo ""

.PHONY: submodule.update-all submodule.update-%

submodule.update-%:
	@echo "--- Running task: $@ ---"
	@echo "Updating submodule: $*..."
	@git submodule update --init --recursive $* 
	@echo "--- Task $@ finished. ---"
	@echo ""

submodule.update-all:
	@echo "--- Running task: $@ ---"
	@echo "Updating all submodules..."
	@git submodule update --init --recursive
	@echo "--- Task $@ finished. ---"
	@echo ""