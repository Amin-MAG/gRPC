package ir.mag.grpc.echo;

import android.app.Activity;
import android.os.AsyncTask;
import android.os.Bundle;
import android.util.Log;

import androidx.appcompat.app.AppCompatActivity;

import com.google.android.material.button.MaterialButton;

import java.io.PrintWriter;
import java.io.StringWriter;
import java.lang.ref.WeakReference;
import java.util.Iterator;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.StatusRuntimeException;

public class MainActivity extends AppCompatActivity {

	private static final String TAG = "MainActivity";

	@Override
	protected void onCreate(Bundle savedInstanceState) {
		super.onCreate(savedInstanceState);
		setContentView(R.layout.activity_main);

		String host = "192.168.1.46";
		int port = 8888;
		ManagedChannel channel = ManagedChannelBuilder.forAddress(host, port).usePlaintext().build();


		MaterialButton materialButton = findViewById(R.id.connect);
		materialButton.setOnClickListener(view -> {
			Log.d(TAG, "onCreate: clicked");
			new GrpcTask(new ListFeaturesRunnable(), channel, this).execute();
		});

	}

	private static class ListFeaturesRunnable implements GrpcRunnable {
		/**
		 * Blocking server-streaming example. Calls listFeatures with a rectangle of interest. Prints
		 * each response feature as it arrives.
		 */
		private void globalCounter(EchoGrpc.EchoBlockingStub blockingStub)
				throws StatusRuntimeException {
			Log.d(TAG, "globalCounter: in the global counter");
			CounterParam request = CounterParam.newBuilder().build();
			Iterator<GlobalCounterResponse> features;
			Log.d(TAG, "globalCounter: set features");
			features = blockingStub.globalCounter(request);

			Log.d(TAG, "globalCounter: starting the while loop");
			while (features.hasNext()) {
				Log.d(TAG, "globalCounter: waiting");
				GlobalCounterResponse feature = features.next();
				Log.d(TAG, "globalCounter: " + feature.getCounter());
			}
		}

		@Override
		public String run(EchoGrpc.EchoBlockingStub blockingStub, EchoGrpc.EchoStub asyncStub) throws Exception {
			Log.d(TAG, "run: going to run the counter");
			globalCounter(blockingStub);
			return "";
		}
	}


	private static class GrpcTask extends AsyncTask<Void, Void, String> {
		private final GrpcRunnable grpcRunnable;
		private final ManagedChannel channel;
		private final WeakReference<MainActivity> activityReference;

		GrpcTask(GrpcRunnable grpcRunnable, ManagedChannel channel, MainActivity activity) {
			this.grpcRunnable = grpcRunnable;
			this.channel = channel;
			this.activityReference = new WeakReference<>(activity);
		}

		@Override
		protected String doInBackground(Void... nothing) {
			try {
				Log.d(TAG, "doInBackground: going to start the connection");
				String logs = grpcRunnable.run(
						EchoGrpc.newBlockingStub(channel),
						EchoGrpc.newStub(channel)
				);
				return "Success!\n" + logs;
			} catch (Exception e) {
				StringWriter sw = new StringWriter();
				PrintWriter pw = new PrintWriter(sw);
				e.printStackTrace(pw);
				pw.flush();
				return "Failed... :\n" + sw;
			}
		}

		@Override
		protected void onPostExecute(String result) {
			MainActivity activity = activityReference.get();
			if (activity == null) {
				return;
			}
		}
	}

	private interface GrpcRunnable {
		/**
		 * Perform a grpcRunnable and return all the logs.
		 */
		String run(EchoGrpc.EchoBlockingStub blockingStub, EchoGrpc.EchoStub asyncStub) throws Exception;
	}

}