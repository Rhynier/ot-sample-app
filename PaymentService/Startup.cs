using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using OpenTelemetry.Resources;
using OpenTelemetry.Trace;

namespace PaymentService
{
    public class Startup
    {
        public Startup(IConfiguration configuration)
        {
            Configuration = configuration;
        }

        public IConfiguration Configuration { get; }

        // This method gets called by the runtime. Use this method to add services to the container.
        public void ConfigureServices(IServiceCollection services)
        {
            services.AddControllers();

            services.AddOpenTelemetryTracing((builder) =>
            {
                builder.AddAspNetCoreInstrumentation();

                switch (Program.SelectedExporter)
                {
                    case Exporter.Console:
                        builder.AddConsoleExporter();
                        break;
                    case Exporter.Jaeger:
                        builder.AddJaegerExporter(o =>
                        {
                            o.AgentHost = Program.JaegerHost;
                        });
                        break;
                    case Exporter.OTC:
                        builder.AddOtlpExporter(o =>
                        {
                            o.Endpoint = new Uri(Program.OtcHost);
                        });
                        break;
                }

                builder.SetResourceBuilder(
                    ResourceBuilder.CreateDefault()
                        .AddService("payments")
                        .AddAttributes(new Dictionary<string, object>
                        {
                            { "k8.cluster", "dev" },
                            { "appdynamics.service", "Payments" },
                            { "service.name", "payments" }
                        }));
            });
        }

        // This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
            }

            app.UseRouting();

            app.UseAuthorization();

            app.UseEndpoints(endpoints =>
            {
                endpoints.MapControllers();
            });
        }
    }
}
