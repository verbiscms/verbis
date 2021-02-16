<!-- =====================
	Settings - Performance
	===================== -->
<template>
	<section>
		<div class="auth-container">
			<div class="row">
				<div class="col-12">
					<header class="header header-with-actions header-margin-large">
						<div class="header-title">
							<h1>Performance</h1>
							<Breadcrumbs></Breadcrumbs>
						</div>
						<div class="header-actions">
							<button class="btn btn-fixed-height btn-orange" @click.prevent="save" :class="{ 'btn-loading' : saving }">
								Update&nbsp;<span class="btn-hide-text-mob">Settings</span>
							</button>
						</div>
					</header>
				</div><!-- /Col -->
			</div><!-- /Row -->
			<!-- Spinner -->
			<div v-show="doingAxios" class="media-spinner spinner-container">
				<div class="spinner spinner-large spinner-grey"></div>
			</div>
			<div v-show="!doingAxios" class="row trans-fade-in-anim">
				<!-- =====================
					Cache Control
					===================== -->
				<div class="col-12">
					<h6 class="margin">Cache control</h6>
					<div class="card card-small-box-shadow card-expand">
						<!-- Cache assets? -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['cache_frontend'] }">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Cache assets?</h4>
										<p>By ticking the box, cache headers will be sent from the Verbis server with the expiration times & extensions set below.</p>
									</div>
									<div class="toggle">
										<input type="checkbox" class="toggle-switch" id="cache-frontend" checked v-model="data['cache_frontend']" :true-value="true" :false-value="false" />
										<label for="cache-frontend"></label>
									</div>
								</div><!-- /Card Header -->
							</template>
						</Collapse><!-- /Cache assets? -->
						<!-- Expiration -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['cache_frontend_request'] ||  errors['cache_frontend_seconds'] }">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Expiration</h4>
										<p>Set an duration in seconds for a duration to set the caching headers.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<!-- Request -->
									<FormGroup label="Request directives">
										<div class="form-select-cont form-input">
											<select class="form-select" v-model.number="data['cache_frontend_request']" @change="getRequestMessage">
												<option selected value="max-age">max-age</option>
												<option value="max-stale">max-stale</option>
												<option value="min-fresh">min-fresh</option>
												<option value="no-cache">no-cache</option>
												<option value="no-store">no-store</option>
												<option value="no-transform">no-store</option>
												<option value="only-if-cached">only-if-cached</option>
											</select>
										</div>
										<p>{{ requestMessage }} Source: <a href="https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control" target="_blank">mozilla.org</a></p>
									</FormGroup><!-- /Request -->
									<!-- Seconds -->
									<FormGroup label="Maximum age*" :error="errors['cache_frontend_seconds']">
										<input class="form-input form-input-white" type="number" v-model.number="data['cache_frontend_seconds']">
										<p>Enter a maximum age (in seconds) to cache the assets. <b>Duration: <span>{{ [data['cache_frontend_seconds'], 'seconds'] | duration('humanize') }}</span></b></p>
									</FormGroup><!-- /Seconds -->
								</div><!-- /Card Body -->
							</template>
						</Collapse><!-- /Expiration -->
						<!-- File extensions -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['cache_frontend_extension'] }">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">File extensions</h4>
										<p>Assets can be cached by inputting file extensions below, if the file extension is not in the list, it will be ignored.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<FormGroup label="File extensions" :error="errors['cache_frontend_extension']">
										<textarea class="form-input form-input-white" rows="12" type="text" v-model="cacheFileExtension"></textarea>
										<p>Enter file extensions types separated by a new line. <b>Note:</b> no need to include the dot (.)</p>
									</FormGroup>
								</div><!-- /Card Body -->
							</template>
						</Collapse><!-- /File extensions -->
					</div><!-- /Card -->
				</div><!-- /Col -->
				<!-- =====================
					Cache Control
					===================== -->
				<div class="col-12">
					<h6 class="margin">Server cache</h6>
					<Alert colour="orange" type="warning">
						<slot><b>Note:</b> The settings below should only be enabled in production.</slot>
					</Alert>
					<div class="card card-small-box-shadow card-expand">
						<!-- Clear Cache -->
						<Collapse :show="false" class="collapse-border-bottom">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Clear cache</h4>
										<p>Click the button to the right to clear the entire server cache.</p>
									</div>
									<div>
										<button class="btn" @click="clearCache" :class="{ 'btn-loading' : clearingCache }">Clear</button>
									</div>
								</div><!-- /Card Header -->
							</template>
						</Collapse><!-- /Cache Renderer -->
						<!-- Cache Renderer -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['cache_server_templates'] }">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Cache templates?</h4>
										<p>By ticking the box, pages will be served from the cache avoiding multiple page renders, drastically improving speed.</p>
									</div>
									<div class="toggle">
										<input type="checkbox" class="toggle-switch" id="cache-database-templates" v-model="data['cache_server_templates']" :true-value="true" :false-value="false" />
										<label for="cache-database-templates"></label>
									</div>
								</div><!-- /Card Header -->
							</template>
						</Collapse><!-- /Cache Renderer -->
						<!-- Cache Field Layouts -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['cache_server_field_layouts'] }">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Cache field layouts?</h4>
										<p>By ticking the box, field layouts will be cached and only read from thedisk once or until updated.</p>
									</div>
									<div class="toggle">
										<input type="checkbox" class="toggle-switch" id="cache-database-field-layouts" v-model="data['cache_server_field_layouts']" :true-value="true" :false-value="false" />
										<label for="cache-database-field-layouts"></label>
									</div>
								</div><!-- /Card Header -->
							</template>
						</Collapse><!-- /Cache Field Layouts -->
					</div><!-- /Card -->
				</div><!-- /Col -->
				<!-- =====================
					Gzip
					===================== -->
				<div class="col-12">
					<h6 class="margin">Gzip compression</h6>
					<div class="card card-small-box-shadow card-expand">
						<!-- Use Gzip compression? -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['gzip'] }">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Use Gzip compression?</h4>
										<p>By ticking the box, the Verbis server will use gzip compression for assets.</p>
									</div>
									<div class="toggle">
										<input type="checkbox" class="toggle-switch" id="gzip" v-model="data['gzip']" checked :true-value="true" :false-value="false" />
										<label for="gzip"></label>
									</div>
								</div><!-- /Card Header -->
							</template>
						</Collapse><!-- /Use Gzip compression? -->
						<!-- Compression -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['gzip_compression'] }">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Compression</h4>
										<p>Set the default compression amount.</p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<FormGroup label="Compression amount">
										<div class="form-select-cont form-input">
											<select class="form-select" v-model.number="data['gzip_compression']">
												<option selected value="best-compression">Best compression</option>
												<option value="best-speed">Best speed</option>
												<option value="default-compression">Default compression</option>
											</select>
										</div>
									</FormGroup>
								</div><!-- /Card Body -->
							</template>
						</Collapse><!-- /Compression -->
						<!-- Use Paths? -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['gzip_use_paths'] }">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Use excluded paths?</h4>
										<p>The default way to set excluded directives for Gzip is to use fiel extensions, if you would like to use paths, check the toggle.</p>
									</div>
									<div class="toggle">
										<input type="checkbox" class="toggle-switch" id="gzip-use-paths" v-model="data['gzip_use_paths']" :true-value="true" :false-value="false" />
										<label for="gzip-use-paths"></label>
									</div>
								</div><!-- /Card Header -->
							</template>
						</Collapse><!-- /Compression -->
						<!-- Excluded File extensions -->
						<Collapse v-if="data['gzip_use_paths']" :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['gzip_excluded_extensions'] }">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Excluded file extensions</h4>
										<p>Set any excluded file extensions to be ignored by gzip compression, such as <code>pdf</code></p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<FormGroup label="Excluded file extensions" :error="errors['gzip_excluded_extensions']">
										<textarea class="form-input form-input-white" rows="12" type="text" v-model="gzipExcludedExtensions"></textarea>
										<p>Enter file extensions types separated by a new line. <b>Note:</b> no need to include the dot (.)</p>
									</FormGroup>
								</div><!-- /Card Body -->
							</template>
						</Collapse><!-- /Excluded File extensions -->
						<!-- Excluded Paths -->
						<Collapse v-else :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['gzip_excluded_paths'] }">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Excluded paths</h4>
										<p>Set any excluded paths to be ignored by the gzip compression such as <code>/assets/pdf</code></p>
									</div>
									<div class="card-controls">
										<i class="feather feather-chevron-down"></i>
									</div>
								</div><!-- /Card Header -->
							</template>
							<template v-slot:body>
								<div class="card-body">
									<FormGroup label="Excluded file extensions" :error="errors['gzip_excluded_paths']">
										<textarea class="form-input form-input-white" rows="12" type="text" v-model="gzipExcludedPaths"></textarea>
										<p>Enter absolute paths separated by a new line.</p>
									</FormGroup>
								</div><!-- /Card Body -->
							</template>
						</Collapse><!-- /Excluded File extensions -->
					</div><!-- /Card -->
				</div><!-- /Col -->
				<!-- =====================
					Minify
					===================== -->
				<div class="col-12">
					<h6 class="margin">Minify</h6>
					<div class="card card-small-box-shadow card-expand">
						<!-- HTML -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['minify_html'] }">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">HTML</h4>
										<p>By ticking the box, the Verbis server will strip the whitespace of HTML files.</p>
									</div>
									<div class="toggle">
										<input type="checkbox" class="toggle-switch" id="minfy-html" v-model="data['minify_html']" :true-value="true" :false-value="false" />
										<label for="minfy-html"></label>
									</div>
								</div><!-- /Card Header -->
							</template>
						</Collapse><!-- /HTML -->
						<!-- JS -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['minify_js'] }">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">Javascript</h4>
										<p>By ticking the box, the Verbis server minify Javascript <code>.js</code> files.</p>
									</div>
									<div class="toggle">
										<input type="checkbox" class="toggle-switch" id="minfy-js" v-model="data['minify_js']" :true-value="true" :false-value="false" />
										<label for="minfy-js"></label>
									</div>
								</div><!-- /Card Header -->
							</template>
						</Collapse><!-- /JS -->
						<!-- CSS -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['minify_css'] }">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">CSS</h4>
										<p>By ticking the box, the Verbis server minify CSS <code>.css</code> files.</p>
									</div>
									<div class="toggle">
										<input type="checkbox" class="toggle-switch" id="minfy-css" v-model="data['minify_css']" :true-value="true" :false-value="false" />
										<label for="minfy-css"></label>
									</div>
								</div><!-- /Card Header -->
							</template>
						</Collapse><!-- /CSS -->
						<!-- SVG -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['minify_svg'] }">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">SVG</h4>
										<p>By ticking the box, the Verbis server minify SVG <code>.svg</code> images.</p>
									</div>
									<div class="toggle">
										<input type="checkbox" class="toggle-switch" id="minfy-svg" v-model="data['minify_svg']" :true-value="true" :false-value="false" />
										<label for="minfy-svg"></label>
									</div>
								</div><!-- /Card Header -->
							</template>
						</Collapse><!-- /SVG -->
						<!-- JSON -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['minify_json'] }">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">JSON</h4>
										<p>By ticking the box, the Verbis server minify SVG <code>.json</code> files.</p>
									</div>
									<div class="toggle">
										<input type="checkbox" class="toggle-switch" id="minfy-json" v-model="data['minify_json']" :true-value="true" :false-value="false" />
										<label for="minfy-json"></label>
									</div>
								</div><!-- /Card Header -->
							</template>
						</Collapse><!-- /JSON -->
						<!-- XML -->
						<Collapse :show="false" class="collapse-border-bottom" :class="{ 'card-expand-error' : errors['minify_xml'] }">
							<template v-slot:header>
								<div class="card-header">
									<div>
										<h4 class="card-title">XML</h4>
										<p>By ticking the box, the Verbis server minify XML <code>.xml</code> files.</p>
									</div>
									<div class="toggle">
										<input type="checkbox" class="toggle-switch" id="minfy-xml" v-model="data['minify_xml']" :true-value="true" :false-value="false" />
										<label for="minfy-xml"></label>
									</div>
								</div><!-- /Card Header -->
							</template>
						</Collapse><!-- /XML -->
					</div><!-- /Card -->
				</div><!-- /Col -->
			</div><!-- /Row -->
		</div><!-- /Container -->
	</section>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Breadcrumbs from "../../components/misc/Breadcrumbs";
import {optionsMixin} from "@/util/options";
import Collapse from "@/components/misc/Collapse";
import FormGroup from "@/components/forms/FormGroup";
import Alert from "@/components/misc/Alert";

export default {
	name: "Performance",
	title: 'Performance Settings',
	mixins: [optionsMixin],
	components: {
		Alert,
		FormGroup,
		Collapse,
		Breadcrumbs
	},
	data: () => ({
		errorMsg: "Fix the errors before saving performance settings.",
		successMsg: "Performance options updated successfully.",
		requestMessage: "",
		clearingCache: false,
	}),
	methods: {
		/*
		 * runAfterGet()
		 * Set teh global request if there is none once axios has finished loading.
		 */
		runAfterGet() {
			const globalRequest = this.data['cache_frontend_request'];
			if (globalRequest === "" || globalRequest === undefined) {
				this.$set(this.data, 'cache_frontend_request', 'max-age');
			}
			const gzipCompression = this.data['gzip_compression'];
			if (gzipCompression === "" || gzipCompression === undefined) {
				this.$set(this.data, 'gzip_compression', 'default-compression');
			}
			this.getRequestMessage();
		},
		/*
		 * getRequestMessage()
		 * Get the request message instructions for the cache request.
		 */
		getRequestMessage() {
			const globalRequest = this.data['cache_frontend_request'];
			switch (globalRequest) {
				case 'max-age': {
					this.requestMessage = 'The maximum amount of time a resource is considered fresh. Unlike Expires, this directive is relative to the time of the request.';
					break;
				}
				case 's-maxage': {
					this.requestMessage = 'Overrides max-age or the Expires header, but only for shared caches (e.g., proxies). Ignored by private caches.';
					break;
				}
				case 'max-stale': {
					this.requestMessage = 'Indicates the client will accept a stale response. An optional value in seconds indicates the upper limit of staleness the client will accept.';
					break;
				}
				case 'min-fresh': {
					this.requestMessage = 'Indicates the client wants a response that will still be fresh for at least the specified number of seconds.';
					break;
				}
				case 'stale-while-revalidate': {
					this.requestMessage = 'Indicates the client will accept a stale response, while asynchronously checking in the background for a fresh one. The seconds value indicates how long the client will accept a stale response.';
					break;
				}
				case 'stale-if-error': {
					this.requestMessage = 'Indicates the client will accept a stale response if the check for a fresh one fails. The seconds value indicates how long the client will accept the stale response after the initial expiration.';
					break;
				}
			}
		},
		/*
	 	 * clearCache()
		 * Clear the server cache.
		 */
		clearCache() {
			this.clearingCache = true;
			this.axios.post("/cache")
				.then(res => {
					this.$noty.success(res.data.message);
				})
				.catch(err => {
					this.handleResponse(err);
				})
				.finally(() => {
					setTimeout(() => {
						this.clearingCache = false;
					}, this.timeoutDelay);
				});
		},
	},
	computed: {
		/*
		 * cacheFileExtension()
		 * Gets returns the cache file extensions with new lines,
		 * set sets the data and transforms into string array.
		 */
		cacheFileExtension: {
			get() {
				const data = this.data['cache_frontend_extensions'];
				return data === undefined ? "" : data.join('\n');
			},
			set(value) {
				this.data['cache_frontend_extensions'] = value.split(/\r\n|\r|\n/);
			}
		},
		/*
		 * gzipExcludedExtensions()
		 * Gets returns the gzip excluded file extensions with new lines,
		 * set sets the data and transforms into string array.
		 */
		gzipExcludedExtensions: {
			get() {
				const data = this.data['gzip_excluded_extensions'];
				return data === undefined ? "" : data.join('\n');
			},
			set(value) {
				this.data['gzip_excluded_file_extensions'] = value.split(/\r\n|\r|\n/);
			}
		},
		/*
		 * gzipExcludedPaths()
		 * Gets returns the gzip excluded paths extensions with new lines,
		 * set sets the data and transforms into string array.
		 */
		gzipExcludedPaths: {
			get() {
				const data = this.data['gzip_excluded_paths'];
				return data === undefined ? "" : data.join('\n');
			},
			set(value) {
				this.data['gzip_excluded_paths'] = value.split(/\r\n|\r|\n/);
			}
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

	// Dummy
	// =========================================================================

</style>