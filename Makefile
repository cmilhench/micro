


SUBDIRS =  service/greetings service/fib service/fac service/people # service/consignments

start:
	@$(MAKE) -j 6 services
.PHONY: start

services: $(SUBDIRS)
.PHONY: services

$(SUBDIRS):
	@for dir in $(SUBDIRS); do \
		$(MAKE) -C $@ start; \
	done
.PHONY: $(SUBDIRS)